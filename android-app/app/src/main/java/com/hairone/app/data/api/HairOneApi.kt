package com.hairone.app.data.api

import com.hairone.app.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface HairOneApi {
    @POST("api/auth/login")
    suspend fun login(@Body request: LoginRequest): Response<AuthResponse>

    @POST("api/auth/register")
    suspend fun register(@Body request: User): Response<Map<String, String>> // Returns {token}

    @GET("api/shops/{id}")
    suspend fun getShop(@Path("id") id: Int): Response<Shop>

    @GET("api/barbers")
    suspend fun getBarbers(@Query("shopId") shopId: Int): Response<List<Barber>>

    @GET("api/bookings")
    suspend fun getBookings(@Query("shopId") shopId: Int): Response<List<Booking>>

    @POST("api/bookings")
    suspend fun createBooking(@Body request: BookingRequest): Response<Booking>

    @POST("api/bookings/{id}/complete")
    suspend fun completeBooking(@Path("id") id: Int, @Body request: CompleteBookingRequest): Response<Unit>
}
