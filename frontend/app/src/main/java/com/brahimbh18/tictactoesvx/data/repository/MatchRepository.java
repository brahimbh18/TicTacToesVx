package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.Match;

import java.util.List;
import java.util.UUID;

public class MatchRepository {

    public Result<Match> createMatch(int boardSize) {
        Match match = new Match(UUID.randomUUID().toString(), boardSize);
        InMemoryStore.MATCHES.put(match.matchId, match);
        return Result.success(match);
    }

    public Result<Match> getMatch(String matchId) {
        Match match = InMemoryStore.MATCHES.get(matchId);
        if (match == null) {
            return Result.failure("Match not found");
        }
        return Result.success(match);
    }

    public Result<Match> makeMove(String matchId, int cellIndex) {
        Match match = InMemoryStore.MATCHES.get(matchId);
        if (match == null) return Result.failure("Match not found");
        if (!"active".equals(match.status)) return Result.failure("Match already finished");
        List<String> board = match.getBoard();
        if (cellIndex < 0 || cellIndex >= board.size() || !board.get(cellIndex).isEmpty()) {
            return Result.failure("Invalid move");
        }
        String playedSymbol = match.nextTurn;
        board.set(cellIndex, playedSymbol);
        if (hasWinner(board, match.boardSize, playedSymbol)) {
            match.status = "finished";
            match.winner = playedSymbol;
            return Result.success(match);
        }
        match.nextTurn = "X".equals(match.nextTurn) ? "O" : "X";
        if (!board.contains("")) {
            match.status = "finished";
            match.winner = "draw";
        }
        return Result.success(match);
    }

    public Result<Match> resign(String matchId) {
        Match match = InMemoryStore.MATCHES.get(matchId);
        if (match == null) return Result.failure("Match not found");
        match.status = "finished";
        match.winner = "X".equals(match.nextTurn) ? "O" : "X";
        return Result.success(match);
    }

    private boolean hasWinner(List<String> board, int size, String symbol) {
        for (int i = 0; i < size; i++) {
            boolean rowWin = true;
            boolean colWin = true;
            for (int j = 0; j < size; j++) {
                if (!symbol.equals(board.get(i * size + j))) {
                    rowWin = false;
                }
                if (!symbol.equals(board.get(j * size + i))) {
                    colWin = false;
                }
            }
            if (rowWin || colWin) return true;
        }

        boolean diagWin = true;
        boolean antiDiagWin = true;
        for (int i = 0; i < size; i++) {
            if (!symbol.equals(board.get(i * size + i))) {
                diagWin = false;
            }
            if (!symbol.equals(board.get(i * size + (size - 1 - i)))) {
                antiDiagWin = false;
            }
        }
        return diagWin || antiDiagWin;
    }
}
