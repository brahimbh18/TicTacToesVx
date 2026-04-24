package com.brahimbh18.tictactoesvx.domain.model;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Match {
    public final String matchId;
    public final int boardSize;
    private final List<String> board;
    public String nextTurn;
    public String winner;
    public String status;

    public Match(String matchId, int boardSize) {
        this.matchId = matchId;
        this.boardSize = boardSize;
        this.board = new ArrayList<>(Collections.nCopies(boardSize * boardSize, ""));
        this.nextTurn = "X";
        this.winner = null;
        this.status = "active";
    }

    public List<String> getBoard() {
        return board;
    }
}
