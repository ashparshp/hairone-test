package com.hairone.app.ui.theme

import android.app.Activity
import androidx.compose.foundation.isSystemInDarkTheme
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.SideEffect
import androidx.compose.ui.graphics.toArgb
import androidx.compose.ui.platform.LocalView
import androidx.core.view.WindowCompat

private val DarkColorScheme = darkColorScheme(
    primary = ColorPrimary,
    background = DarkBackground,
    surface = DarkCard,
    onPrimary = DarkText,
    onBackground = DarkText,
    onSurface = DarkText,
    secondary = DarkIconActive,
    outline = DarkBorder
)

private val LightColorScheme = lightColorScheme(
    primary = ColorPrimary,
    background = LightBackground,
    surface = LightCard,
    onPrimary = LightText,
    onBackground = LightText,
    onSurface = LightText,
    secondary = LightIconActive,
    outline = LightBorder
)

@Composable
fun HairOneTheme(
    darkTheme: Boolean = isSystemInDarkTheme(),
    content: @Composable () -> Unit
) {
    val colorScheme = if (darkTheme) DarkColorScheme else LightColorScheme
    val view = LocalView.current
    if (!view.isInEditMode) {
        SideEffect {
            val window = (view.context as Activity).window
            window.statusBarColor = colorScheme.background.toArgb()
            WindowCompat.getInsetsController(window, view).isAppearanceLightStatusBars = !darkTheme
        }
    }

    MaterialTheme(
        colorScheme = colorScheme,
        content = content
    )
}
