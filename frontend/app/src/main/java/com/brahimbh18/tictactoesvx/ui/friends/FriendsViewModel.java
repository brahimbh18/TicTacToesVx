package com.brahimbh18.tictactoesvx.ui.friends;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.Friend;
import com.brahimbh18.tictactoesvx.domain.model.FriendRequest;
import com.brahimbh18.tictactoesvx.domain.model.User;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

import java.util.List;

public class FriendsViewModel extends ViewModel {
    private final MutableLiveData<UiState<List<User>>> searchState = new MutableLiveData<>(UiState.idle());
    private final MutableLiveData<UiState<List<FriendRequest>>> requestsState = new MutableLiveData<>(UiState.idle());
    private final MutableLiveData<UiState<List<Friend>>> friendsState = new MutableLiveData<>(UiState.idle());
    private final MutableLiveData<UiState<String>> actionState = new MutableLiveData<>(UiState.idle());

    public LiveData<UiState<List<User>>> getSearchState() { return searchState; }
    public LiveData<UiState<List<FriendRequest>>> getRequestsState() { return requestsState; }
    public LiveData<UiState<List<Friend>>> getFriendsState() { return friendsState; }
    public LiveData<UiState<String>> getActionState() { return actionState; }

    public void searchUsers(String query) {
        Result<List<User>> result = ServiceLocator.friendRepository().searchUsers(query);
        searchState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
    }

    public void sendFriendRequest(String username) {
        Result<String> result = ServiceLocator.friendRepository().sendFriendRequest(username);
        actionState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
    }

    public void loadIncomingRequests() {
        Result<List<FriendRequest>> result = ServiceLocator.friendRepository().getIncomingRequests();
        requestsState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
    }

    public void accept(String requestId) {
        Result<String> result = ServiceLocator.friendRepository().acceptRequest(requestId);
        actionState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
        loadIncomingRequests();
        loadFriends();
    }

    public void decline(String requestId) {
        Result<String> result = ServiceLocator.friendRepository().declineRequest(requestId);
        actionState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
        loadIncomingRequests();
    }

    public void loadFriends() {
        Result<List<Friend>> result = ServiceLocator.friendRepository().getFriends();
        friendsState.setValue(result.isSuccess() ? UiState.success(result.getData()) : UiState.error(result.getError()));
    }
}
