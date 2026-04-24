package com.brahimbh18.tictactoesvx.ui.auth.login;

import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.ui.auth.register.RegisterActivity;
import com.brahimbh18.tictactoesvx.ui.home.HomeActivity;

public class LoginActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login);

        LoginViewModel viewModel = new ViewModelProvider(this).get(LoginViewModel.class);
        EditText username = findViewById(R.id.etUsername);
        EditText password = findViewById(R.id.etPassword);
        Button loginBtn = findViewById(R.id.btnLogin);
        TextView registerLink = findViewById(R.id.tvRegisterLink);

        loginBtn.setOnClickListener(v -> viewModel.login(username.getText().toString().trim(), password.getText().toString().trim()));
        registerLink.setOnClickListener(v -> startActivity(new Intent(this, RegisterActivity.class)));

        viewModel.getLoginState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS) {
                startActivity(new Intent(this, HomeActivity.class));
                finish();
            } else if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.message, Toast.LENGTH_SHORT).show();
            }
        });
    }
}
