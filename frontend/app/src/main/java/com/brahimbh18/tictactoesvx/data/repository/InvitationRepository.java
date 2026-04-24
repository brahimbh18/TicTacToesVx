package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.Invitation;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

public class InvitationRepository {
    public Result<List<Invitation>> getIncomingInvitations() {
        return Result.success(new ArrayList<>(InMemoryStore.INVITATIONS));
    }

    public Result<Invitation> seedInvitation(String fromUsername, int boardSize) {
        Invitation invitation = new Invitation(UUID.randomUUID().toString(), fromUsername, boardSize);
        InMemoryStore.INVITATIONS.add(invitation);
        return Result.success(invitation);
    }

    public Result<String> acceptInvitation(String invitationId) {
        InMemoryStore.INVITATIONS.removeIf(i -> i.invitationId.equals(invitationId));
        return Result.success("match_" + invitationId);
    }

    public Result<String> declineInvitation(String invitationId) {
        InMemoryStore.INVITATIONS.removeIf(i -> i.invitationId.equals(invitationId));
        return Result.success("Declined");
    }
}
