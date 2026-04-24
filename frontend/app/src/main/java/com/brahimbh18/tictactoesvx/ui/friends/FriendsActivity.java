package com.brahimbh18.tictactoesvx.ui.friends;

import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.domain.model.FriendRequest;
import com.brahimbh18.tictactoesvx.domain.model.User;

import java.util.List;

public class FriendsActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_friends);

        FriendsViewModel viewModel = new ViewModelProvider(this).get(FriendsViewModel.class);
        EditText query = findViewById(R.id.etSearchUsername);
        TextView searchOut = findViewById(R.id.tvSearchResults);
        TextView requestsOut = findViewById(R.id.tvIncomingRequests);
        TextView friendsOut = findViewById(R.id.tvFriendsList);
        EditText requestIdInput = findViewById(R.id.etRequestId);

        findViewById(R.id.btnSearchUsers).setOnClickListener(v -> viewModel.searchUsers(query.getText().toString().trim()));
        findViewById(R.id.btnSendRequest).setOnClickListener(v -> viewModel.sendFriendRequest(query.getText().toString().trim()));
        findViewById(R.id.btnLoadRequests).setOnClickListener(v -> viewModel.loadIncomingRequests());
        findViewById(R.id.btnAcceptRequest).setOnClickListener(v -> viewModel.accept(requestIdInput.getText().toString().trim()));
        findViewById(R.id.btnDeclineRequest).setOnClickListener(v -> viewModel.decline(requestIdInput.getText().toString().trim()));
        findViewById(R.id.btnLoadFriends).setOnClickListener(v -> viewModel.loadFriends());

        viewModel.getSearchState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS) {
                searchOut.setText(renderUsers(state.data));
            }
        });

        viewModel.getRequestsState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS) {
                requestsOut.setText(renderRequests(state.data));
            }
        });

        viewModel.getFriendsState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS && state.data != null) {
                StringBuilder sb = new StringBuilder();
                for (com.brahimbh18.tictactoesvx.domain.model.Friend friend : state.data) {
                    sb.append(friend.username).append("\n");
                }
                friendsOut.setText(sb.length() == 0 ? getString(R.string.empty_state) : sb.toString());
            }
        });

        viewModel.getActionState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS || state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS ? state.data : state.message, Toast.LENGTH_SHORT).show();
            }
        });
    }

    private String renderUsers(List<User> users) {
        if (users == null || users.isEmpty()) return getString(R.string.empty_state);
        StringBuilder sb = new StringBuilder();
        for (User user : users) {
            sb.append(user.username).append("\n");
        }
        return sb.toString();
    }

    private String renderRequests(List<FriendRequest> requests) {
        if (requests == null || requests.isEmpty()) return getString(R.string.empty_state);
        StringBuilder sb = new StringBuilder();
        for (FriendRequest request : requests) {
            String requestId = request.requestId == null ? "unknown" : request.requestId;
            String fromUsername = request.fromUsername == null ? "unknown" : request.fromUsername;
            sb.append(requestId, 0, Math.min(8, requestId.length()))
                    .append("...").append(" from ").append(fromUsername).append("\n");
        }
        return sb.toString();
    }
}
