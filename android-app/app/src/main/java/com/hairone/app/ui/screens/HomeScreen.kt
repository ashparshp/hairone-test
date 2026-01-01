package com.hairone.app.ui.screens

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import com.hairone.app.data.model.Shop
import com.hairone.app.ui.components.ShopCard
import com.hairone.app.ui.theme.Spacing

// Simplified logic - usually fetched via ViewModel
@Composable
fun HomeScreen(
    onShopClick: (Int) -> Unit
) {
    // Mock Data for UI Structure
    val shops = remember {
        listOf(
            Shop(1, 1, "Classic Cuts", "123 Main St", null, 4.8, 120, "male"),
            Shop(2, 2, "Style Salon", "456 Oak Ave", null, 4.5, 85, "unisex")
        )
    }

    LazyColumn(
        modifier = Modifier
            .fillMaxSize()
            .padding(Spacing.screenPadding)
    ) {
        item {
            Text("Nearby Shops", modifier = Modifier.padding(bottom = Spacing.md))
        }
        items(shops.size) { index ->
            ShopCard(shop = shops[index], onClick = { onShopClick(shops[index].id) })
        }
    }
}
