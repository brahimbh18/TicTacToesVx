package com.brahimbh18.tictactoesvx.core.network;

import com.brahimbh18.tictactoesvx.core.storage.TokenStorage;

import java.net.HttpURLConnection;

public class AuthInterceptor {
    private final TokenStorage tokenStorage;

    public AuthInterceptor(TokenStorage tokenStorage) {
        this.tokenStorage = tokenStorage;
    }

    public void apply(HttpURLConnection connection) {
        String token = tokenStorage.getToken();
        if (token != null && !token.isEmpty()) {
            connection.setRequestProperty("Authorization", "Bearer " + token);
        }
    }
}
