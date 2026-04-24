package com.brahimbh18.tictactoesvx.ui.game;

import android.os.Bundle;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.Spinner;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;
import androidx.recyclerview.widget.GridLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.domain.model.AIDifficulty;
import com.brahimbh18.tictactoesvx.domain.model.GameMode;

import java.util.Arrays;

public class GameActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_game);

        GameViewModel viewModel = new ViewModelProvider(this).get(GameViewModel.class);
        Spinner modeSpinner = findViewById(R.id.spinnerMode);
        Spinner difficultySpinner = findViewById(R.id.spinnerDifficulty);
        Spinner boardSpinner = findViewById(R.id.spinnerBoardSize);
        TextView status = findViewById(R.id.tvGameStatus);
        RecyclerView boardRecycler = findViewById(R.id.rvBoard);
        Button startBtn = findViewById(R.id.btnStartMatch);
        Button resignBtn = findViewById(R.id.btnResign);

        modeSpinner.setAdapter(new ArrayAdapter<>(this, android.R.layout.simple_spinner_dropdown_item, Arrays.asList("LOCAL", "AI")));
        difficultySpinner.setAdapter(new ArrayAdapter<>(this, android.R.layout.simple_spinner_dropdown_item, Arrays.asList("EASY", "MEDIUM", "HARD")));
        boardSpinner.setAdapter(new ArrayAdapter<>(this, android.R.layout.simple_spinner_dropdown_item, Arrays.asList("3", "4", "5")));

        BoardAdapter adapter = new BoardAdapter(viewModel::makeMove);
        boardRecycler.setAdapter(adapter);

        startBtn.setOnClickListener(v -> {
            viewModel.setMode("AI".equals(modeSpinner.getSelectedItem().toString()) ? GameMode.AI : GameMode.LOCAL);
            viewModel.setDifficulty(AIDifficulty.valueOf(difficultySpinner.getSelectedItem().toString()));
            viewModel.setBoardSize(Integer.parseInt(boardSpinner.getSelectedItem().toString()));
            viewModel.createMatch();
        });

        resignBtn.setOnClickListener(v -> viewModel.resign());

        viewModel.getMatchState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS && state.data != null) {
                boardRecycler.setLayoutManager(new GridLayoutManager(this, state.data.boardSize));
                adapter.submit(state.data.getBoard());
                status.setText(getString(R.string.game_status, state.data.status, state.data.nextTurn, state.data.winner == null ? "-" : state.data.winner));
            } else if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.message, Toast.LENGTH_SHORT).show();
            }
        });
    }
}
