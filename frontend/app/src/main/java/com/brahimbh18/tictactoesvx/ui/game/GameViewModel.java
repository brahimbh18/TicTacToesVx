package com.brahimbh18.tictactoesvx.ui.game;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.AIDifficulty;
import com.brahimbh18.tictactoesvx.domain.model.GameMode;
import com.brahimbh18.tictactoesvx.domain.model.Match;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

public class GameViewModel extends ViewModel {
    private final MutableLiveData<UiState<Match>> matchState = new MutableLiveData<>(UiState.idle());
    private final MutableLiveData<GameMode> mode = new MutableLiveData<>(GameMode.LOCAL);
    private final MutableLiveData<AIDifficulty> difficulty = new MutableLiveData<>(AIDifficulty.EASY);
    private final MutableLiveData<Integer> boardSize = new MutableLiveData<>(3);

    private String currentMatchId;

    public LiveData<UiState<Match>> getMatchState() { return matchState; }
    public LiveData<GameMode> getMode() { return mode; }
    public LiveData<AIDifficulty> getDifficulty() { return difficulty; }
    public LiveData<Integer> getBoardSize() { return boardSize; }

    public void setMode(GameMode gameMode) { mode.setValue(gameMode); }
    public void setDifficulty(AIDifficulty aiDifficulty) { difficulty.setValue(aiDifficulty); }
    public void setBoardSize(int size) { boardSize.setValue(size); }

    public void createMatch() {
        Result<Match> result = ServiceLocator.matchRepository().createMatch(boardSize.getValue() == null ? 3 : boardSize.getValue());
        if (result.isSuccess()) {
            currentMatchId = result.getData().matchId;
            matchState.setValue(UiState.success(result.getData()));
        } else {
            matchState.setValue(UiState.error(result.getError()));
        }
    }

    public void makeMove(int index) {
        if (currentMatchId == null) {
            matchState.setValue(UiState.error("No active match"));
            return;
        }
        Result<Match> result = ServiceLocator.matchRepository().makeMove(currentMatchId, index);
        if (result.isSuccess()) {
            matchState.setValue(UiState.success(result.getData()));
        } else {
            matchState.setValue(UiState.error(result.getError()));
        }
    }

    public void resign() {
        if (currentMatchId == null) return;
        Result<Match> result = ServiceLocator.matchRepository().resign(currentMatchId);
        if (result.isSuccess()) {
            matchState.setValue(UiState.success(result.getData()));
        }
    }
}
