package com.hairone.app.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.hairone.app.data.api.HairOneApi
import com.hairone.app.data.model.BookingRequest
import com.hairone.app.data.model.CompleteBookingRequest
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

sealed class BookingState {
    object Idle : BookingState()
    object Loading : BookingState()
    object Success : BookingState()
    data class Error(val message: String) : BookingState()
}

@HiltViewModel
class BookingViewModel @Inject constructor(
    private val api: HairOneApi
) : ViewModel() {

    private val _state = MutableStateFlow<BookingState>(BookingState.Idle)
    val state: StateFlow<BookingState> = _state

    fun createBooking(request: BookingRequest) {
        viewModelScope.launch {
            _state.value = BookingState.Loading
            try {
                val res = api.createBooking(request)
                if (res.isSuccessful) _state.value = BookingState.Success
                else _state.value = BookingState.Error("Booking failed")
            } catch (e: Exception) {
                _state.value = BookingState.Error(e.message ?: "Error")
            }
        }
    }

    fun completeBooking(id: Int, pin: String) {
        viewModelScope.launch {
            _state.value = BookingState.Loading
            try {
                val res = api.completeBooking(id, CompleteBookingRequest(pin))
                if (res.isSuccessful) _state.value = BookingState.Success
                else _state.value = BookingState.Error("Invalid PIN")
            } catch (e: Exception) {
                 _state.value = BookingState.Error(e.message ?: "Error")
            }
        }
    }
}
