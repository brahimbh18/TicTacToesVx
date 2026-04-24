package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.core.storage.TokenStorage;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.User;

public class UserRepository {
    private final TokenStorage tokenStorage;

    public UserRepository(TokenStorage tokenStorage) {
        this.tokenStorage = tokenStorage;
    }

    public Result<User> getCurrentUser() {
        String token = tokenStorage.getToken();
        if (token == null) {
            return Result.failure("Unauthorized");
        }
        User user = InMemoryStore.userFromToken(token);
        if (user == null) {
            return Result.failure("Unauthorized");
        }
        return Result.success(user);
    }
}
