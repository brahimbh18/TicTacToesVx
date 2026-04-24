package com.brahimbh18.tictactoesvx.ui.invitations;

import android.os.Bundle;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.lifecycle.ViewModelProvider;

import com.brahimbh18.tictactoesvx.R;
import com.brahimbh18.tictactoesvx.domain.model.Invitation;

import java.util.List;

public class InvitationsActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_invitations);

        InvitationsViewModel viewModel = new ViewModelProvider(this).get(InvitationsViewModel.class);
        EditText invitationIdInput = findViewById(R.id.etInvitationId);
        TextView invitationsOut = findViewById(R.id.tvInvitations);

        findViewById(R.id.btnLoadInvitations).setOnClickListener(v -> viewModel.loadIncoming());
        findViewById(R.id.btnAcceptInvitation).setOnClickListener(v -> viewModel.accept(invitationIdInput.getText().toString().trim()));
        findViewById(R.id.btnDeclineInvitation).setOnClickListener(v -> viewModel.decline(invitationIdInput.getText().toString().trim()));

        viewModel.getInvitationsState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS) {
                invitationsOut.setText(renderInvitations(state.data));
            }
        });

        viewModel.getActionState().observe(this, state -> {
            if (state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS || state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.ERROR) {
                Toast.makeText(this, state.status == com.brahimbh18.tictactoesvx.ui.common.UiState.Status.SUCCESS ? state.data : state.message, Toast.LENGTH_SHORT).show();
            }
        });

        viewModel.loadIncoming();
    }

    private String renderInvitations(List<Invitation> invitations) {
        if (invitations == null || invitations.isEmpty()) return getString(R.string.empty_state);
        StringBuilder sb = new StringBuilder();
        for (Invitation invitation : invitations) {
            String invitationId = invitation.invitationId == null ? "unknown" : invitation.invitationId;
            String fromUsername = invitation.fromUsername == null ? "unknown" : invitation.fromUsername;
            sb.append(invitationId, 0, Math.min(8, invitationId.length()))
                    .append("...")
                    .append(" from ").append(fromUsername)
                    .append(" board ").append(invitation.boardSize)
                    .append("\n");
        }
        return sb.toString();
    }
}
