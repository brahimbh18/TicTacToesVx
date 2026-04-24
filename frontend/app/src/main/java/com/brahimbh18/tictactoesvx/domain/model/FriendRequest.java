package com.brahimbh18.tictactoesvx.domain.model;

public class FriendRequest {
    public final String requestId;
    public final String fromUsername;

    public FriendRequest(String requestId, String fromUsername) {
        this.requestId = requestId;
        this.fromUsername = fromUsername;
    }
}
