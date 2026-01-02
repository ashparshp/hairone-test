const cron = require('node-cron');
const Booking = require('../models/Booking');
const User = require('../models/User');
const SystemConfig = require('../models/SystemConfig');
const { getISTTime } = require('../utils/dateUtils');
const { timeToMinutes } = require('../utils/scheduleUtils');

/**
 * AUTO CANCEL / MISSED BOOKING JOB
 * Runs every 30 minutes to check for bookings that have passed their end time
 * without being completed.
 */
const runAutoCancelJob = () => {
    // Schedule: Every 30 minutes
    cron.schedule('*/30 * * * *', async () => {
        console.log("--- RUNNING AUTO-MISSED BOOKING JOB ---");
        try {
            const { date: istDate, minutes: istMinutes } = getISTTime();

            // Get Config for Limits
            const config = await SystemConfig.findOne({ key: 'global' });
            const limit = (config && config.yearlyCancellationLimit) ? config.yearlyCancellationLimit : 12;

            // Find bookings that are 'upcoming' (or pending)
            // We need to fetch them and check the time manually because time is stored as string "HH:mm"
            // Optimization: Filter by date <= istDate
            const bookings = await Booking.find({
                status: { $in: ['upcoming', 'pending'] },
                date: { $lte: istDate }
            });

            let count = 0;
            for (const b of bookings) {
                // If date is in the past (yesterday etc), it's definitely missed
                let isMissed = false;
                if (b.date < istDate) {
                    isMissed = true;
                } else if (b.date === istDate) {
                    // If date is today, check if end time has passed
                    const endTimeMin = timeToMinutes(b.endTime);
                    if (endTimeMin < istMinutes) {
                        isMissed = true;
                    }
                }

                if (isMissed) {
                    // Update Booking
                    b.status = 'missed';
                    await b.save();
                    count++;

                    // Update User Stats Atomically
                    if (b.userId) {
                        // Increment noShowCount and fetch updated user to check flags
                        const user = await User.findByIdAndUpdate(
                            b.userId,
                            { $inc: { noShowCount: 1 } },
                            { new: true }
                        );

                        if (user) {
                            const totalIncidents = (user.cancellationCount || 0) + user.noShowCount;
                            if (totalIncidents > limit && !user.isFlagged) {
                                await User.findByIdAndUpdate(user._id, { isFlagged: true });
                            }
                        }
                    }
                }
            }

            if (count > 0) {
                console.log(`Auto-Missed Job: Marked ${count} bookings as missed.`);
            }

        } catch (error) {
            console.error("Error in Auto-Missed Job:", error);
        }
    });
};

module.exports = runAutoCancelJob;
