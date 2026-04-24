package com.brahimbh18.tictactoesvx.ui.home;

import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.ui.auth.login.LoginActivity;
import com.brahimbh18.tictactoesvx.ui.friends.FriendsActivity;
import com.brahimbh18.tictactoesvx.ui.game.GameActivity;
import com.brahimbh18.tictactoesvx.ui.invitations.InvitationsActivity;

public class HomeActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_home);

        HomeViewModel viewModel = new ViewModelProvider(this).get(HomeViewModel.class);
        TextView welcome = findViewById(R.id.tvWelcome);

        findViewById(R.id.btnOpenGame).setOnClickListener(v -> startActivity(new Intent(this, GameActivity.class)));
        findViewById(R.id.btnOpenFriends).setOnClickListener(v -> startActivity(new Intent(this, FriendsActivity.class)));
        findViewById(R.id.btnOpenInvitations).setOnClickListener(v -> startActivity(new Intent(this, InvitationsActivity.class)));
        Button logout = findViewById(R.id.btnLogout);
        logout.setOnClickListener(v -> {
            viewModel.logout();
            startActivity(new Intent(this, LoginActivity.class));
            finish();
        });

        viewModel.getMeState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS && state.data != null) {
                welcome.setText(getString(R.string.welcome_user, state.data.username));
            } else if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.message, Toast.LENGTH_SHORT).show();
                startActivity(new Intent(this, LoginActivity.class));
                finish();
            }
        });

        viewModel.loadMe();
    }
}
