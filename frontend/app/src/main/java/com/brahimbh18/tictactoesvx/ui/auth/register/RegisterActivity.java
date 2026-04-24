package com.brahimbh18.tictactoesvx.ui.auth.register;

import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.ui.auth.login.LoginActivity;
import com.brahimbh18.tictactoesvx.ui.home.HomeActivity;

public class RegisterActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_register);

        RegisterViewModel viewModel = new ViewModelProvider(this).get(RegisterViewModel.class);
        EditText username = findViewById(R.id.etUsername);
        EditText password = findViewById(R.id.etPassword);
        Button registerBtn = findViewById(R.id.btnRegister);
        TextView loginLink = findViewById(R.id.tvLoginLink);

        registerBtn.setOnClickListener(v -> viewModel.register(username.getText().toString().trim(), password.getText().toString()));
        loginLink.setOnClickListener(v -> startActivity(new Intent(this, LoginActivity.class)));

        viewModel.getRegisterState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS) {
                startActivity(new Intent(this, HomeActivity.class));
                finish();
            } else if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.message, Toast.LENGTH_SHORT).show();
            }
        });
    }
}
