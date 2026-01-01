package com.hairone.app.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.hairone.app.ui.components.AppInput
import com.hairone.app.ui.components.PrimaryButton
import com.hairone.app.ui.theme.Spacing
import com.hairone.app.viewmodel.AuthState
import com.hairone.app.viewmodel.AuthViewModel

@Composable
fun LoginScreen(
    onLoginSuccess: () -> Unit,
    viewModel: AuthViewModel = hiltViewModel()
) {
    var phone by remember { mutableStateOf("") }
    var password by remember { mutableStateOf("") }
    val authState by viewModel.authState.collectAsState()

    LaunchedEffect(authState) {
        if (authState is AuthState.Success) {
            onLoginSuccess()
        }
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(Spacing.screenPadding),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Text("HairOne", style = MaterialTheme.typography.displayMedium, color = MaterialTheme.colorScheme.primary)

        Spacer(modifier = Modifier.height(Spacing.xl))

        AppInput(value = phone, onValueChange = { phone = it }, label = "Phone Number")
        Spacer(modifier = Modifier.height(Spacing.md))
        AppInput(value = password, onValueChange = { password = it }, label = "Password")

        if (authState is AuthState.Error) {
            Spacer(modifier = Modifier.height(Spacing.md))
            Text((authState as AuthState.Error).message, color = MaterialTheme.colorScheme.error)
        }

        Spacer(modifier = Modifier.height(Spacing.lg))

        PrimaryButton(
            text = if (authState is AuthState.Loading) "Logging in..." else "Login",
            onClick = { viewModel.login(phone, password) },
            enabled = authState !is AuthState.Loading
        )
    }
}
