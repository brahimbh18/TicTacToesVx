package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.core.storage.TokenStorage;
import com.brahimbh18.tictactoesvx.core.util.Result;

import java.util.UUID;

public class AuthRepository {
    private final TokenStorage tokenStorage;

    public AuthRepository(TokenStorage tokenStorage) {
        this.tokenStorage = tokenStorage;
    }

    public Result<String> register(String username, String password) {
        if (username == null || username.trim().isEmpty() || password == null || password.length() < 8) {
            return Result.failure("Password must be at least 8 characters");
        }
        if (InMemoryStore.USERS.containsKey(username)) {
            return Result.failure("Username already exists");
        }
        InMemoryStore.USERS.put(username, password);
        InMemoryStore.USER_IDS.put(username, UUID.randomUUID().toString());
        String token = InMemoryStore.issueToken(username);
        tokenStorage.saveToken(token);
        return Result.success(token);
    }

    public Result<String> login(String username, String password) {
        String stored = InMemoryStore.USERS.get(username);
        if (stored == null || !stored.equals(password)) {
            return Result.failure("Invalid credentials");
        }
        String token = InMemoryStore.issueToken(username);
        tokenStorage.saveToken(token);
        return Result.success(token);
    }

    public void logout() {
        tokenStorage.clear();
    }
}
