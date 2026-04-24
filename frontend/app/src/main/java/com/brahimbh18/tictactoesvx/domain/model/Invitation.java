package com.brahimbh18.tictactoesvx.domain.model;

public class Invitation {
    public final String invitationId;
    public final String fromUsername;
    public final int boardSize;

    public Invitation(String invitationId, String fromUsername, int boardSize) {
        this.invitationId = invitationId;
        this.fromUsername = fromUsername;
        this.boardSize = boardSize;
    }
}
