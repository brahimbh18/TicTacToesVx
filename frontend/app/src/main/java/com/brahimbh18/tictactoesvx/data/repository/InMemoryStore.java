package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.domain.model.Friend;
import com.brahimbh18.tictactoesvx.domain.model.Invitation;
import com.brahimbh18.tictactoesvx.domain.model.Match;
import com.brahimbh18.tictactoesvx.domain.model.User;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

final class InMemoryStore {
    static final Map<String, String> USERS = new HashMap<>();
    static final Map<String, String> USER_IDS = new HashMap<>();
    static final Map<String, String> TOKENS = new HashMap<>();
    static final List<Friend> FRIENDS = new ArrayList<>();
    static final List<FriendRequest> REQUESTS = new ArrayList<>();
    static final List<Invitation> INVITATIONS = new ArrayList<>();
    static final Map<String, Match> MATCHES = new HashMap<>();

    static {
        USERS.put("demo", "demo123");
        USER_IDS.put("demo", UUID.randomUUID().toString());
    }

    static String issueToken(String username) {
        String token = "token_" + username + "_" + System.currentTimeMillis();
        TOKENS.put(token, username);
        return token;
    }

    static User userFromToken(String token) {
        String username = TOKENS.get(token);
        if (username == null) return null;
        return new User(USER_IDS.get(username), username);
    }

    static final class FriendRequest {
        final String requestId;
        final String fromUsername;

        FriendRequest(String fromUsername) {
            this.requestId = UUID.randomUUID().toString();
            this.fromUsername = fromUsername;
        }
    }

    private InMemoryStore() {}
}
