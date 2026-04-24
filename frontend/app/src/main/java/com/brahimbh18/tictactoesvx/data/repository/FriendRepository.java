package com.brahimbh18.tictactoesvx.data.repository;

import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.Friend;
import com.brahimbh18.tictactoesvx.domain.model.FriendRequest;
import com.brahimbh18.tictactoesvx.domain.model.User;

import java.util.ArrayList;
import java.util.List;

public class FriendRepository {

    public Result<List<User>> searchUsers(String query) {
        List<User> result = new ArrayList<>();
        for (String username : InMemoryStore.USERS.keySet()) {
            if (username.toLowerCase().contains(query.toLowerCase())) {
                result.add(new User(InMemoryStore.USER_IDS.get(username), username));
            }
        }
        return Result.success(result);
    }

    public Result<String> sendFriendRequest(String toUsername) {
        InMemoryStore.REQUESTS.add(new InMemoryStore.FriendRequest(toUsername));
        return Result.success("Request sent");
    }

    public Result<List<FriendRequest>> getIncomingRequests() {
        List<FriendRequest> requests = new ArrayList<>();
        for (InMemoryStore.FriendRequest request : InMemoryStore.REQUESTS) {
            requests.add(new FriendRequest(request.requestId, request.fromUsername));
        }
        return Result.success(requests);
    }

    public Result<String> acceptRequest(String requestId) {
        InMemoryStore.FriendRequest target = null;
        for (InMemoryStore.FriendRequest req : InMemoryStore.REQUESTS) {
            if (req.requestId.equals(requestId)) {
                target = req;
                break;
            }
        }
        if (target == null) return Result.failure("Request not found");
        InMemoryStore.REQUESTS.remove(target);
        InMemoryStore.FRIENDS.add(new Friend(InMemoryStore.USER_IDS.get(target.fromUsername), target.fromUsername));
        return Result.success("Accepted");
    }

    public Result<String> declineRequest(String requestId) {
        InMemoryStore.REQUESTS.removeIf(req -> req.requestId.equals(requestId));
        return Result.success("Declined");
    }

    public Result<List<Friend>> getFriends() {
        return Result.success(new ArrayList<>(InMemoryStore.FRIENDS));
    }
}
