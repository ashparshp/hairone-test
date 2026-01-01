package com.hairone.app.data.model

import com.squareup.moshi.Json

data class User(
    @Json(name = "id") val id: Int,
    @Json(name = "phone") val phone: String,
    @Json(name = "name") val name: String,
    @Json(name = "avatar") val avatar: String? = null,
    @Json(name = "role") val role: String
)

data class Shop(
    @Json(name = "id") val id: Int,
    @Json(name = "ownerId") val ownerId: Int,
    @Json(name = "name") val name: String,
    @Json(name = "address") val address: String,
    @Json(name = "image") val image: String? = null,
    @Json(name = "rating") val rating: Double,
    @Json(name = "reviewCount") val reviewCount: Int,
    @Json(name = "type") val type: String
)

data class Barber(
    @Json(name = "id") val id: Int,
    @Json(name = "shopId") val shopId: Int,
    @Json(name = "name") val name: String,
    @Json(name = "avatar") val avatar: String? = null
)

data class Booking(
    @Json(name = "id") val id: Int,
    @Json(name = "shopId") val shopId: Int,
    @Json(name = "date") val date: String,
    @Json(name = "startTime") val startTime: String,
    @Json(name = "status") val status: String,
    @Json(name = "totalPrice") val totalPrice: Double
)

data class AuthResponse(
    @Json(name = "token") val token: String,
    @Json(name = "user") val user: User
)

data class LoginRequest(
    @Json(name = "phone") val phone: String,
    @Json(name = "password") val password: String
)

data class BookingRequest(
    @Json(name = "shopId") val shopId: Int,
    @Json(name = "barberId") val barberId: Int,
    @Json(name = "serviceNames") val serviceNames: List<String>,
    @Json(name = "totalPrice") val totalPrice: Double,
    @Json(name = "totalDuration") val totalDuration: Int,
    @Json(name = "date") val date: String,
    @Json(name = "startTime") val startTime: String,
    @Json(name = "endTime") val endTime: String
)

data class CompleteBookingRequest(
    @Json(name = "bookingKey") val bookingKey: String
)
