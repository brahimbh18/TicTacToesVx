package com.brahimbh18.tictactoesvx.ui.invitations;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.Invitation;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

import java.util.List;

public class InvitationsViewModel extends ViewModel {
    private final MutableLiveData<UiState<List<Invitation>>> invitationsState = new MutableLiveData<>(UiState.idle());
    private final MutableLiveData<UiState<String>> actionState = new MutableLiveData<>(UiState.idle());

    public LiveData<UiState<List<Invitation>>> getInvitationsState() { return invitationsState; }
    public LiveData<UiState<String>> getActionState() { return actionState; }

    public void loadIncoming() {
        Result<List<Invitation>> result = ServiceLocator.invitationRepository().getIncomingInvitations();
        invitationsState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
    }

    public void accept(String invitationId) {
        Result<String> result = ServiceLocator.invitationRepository().acceptInvitation(invitationId);
        actionState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
        loadIncoming();
    }

    public void decline(String invitationId) {
        Result<String> result = ServiceLocator.invitationRepository().declineInvitation(invitationId);
        actionState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
        loadIncoming();
    }
}
