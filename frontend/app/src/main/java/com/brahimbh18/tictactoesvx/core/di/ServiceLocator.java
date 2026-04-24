package com.brahimbh18.tictactoesvx.core.di;

import android.content.Context;

import com.brahimbh18.tictactoesvx.core.network.ApiClient;
import com.brahimbh18.tictactoesvx.core.network.ApiService;
import com.brahimbh18.tictactoesvx.core.storage.TokenStorage;
import com.brahimbh18.tictactoesvx.data.repository.AuthRepository;
import com.brahimbh18.tictactoesvx.data.repository.FriendRepository;
import com.brahimbh18.tictactoesvx.data.repository.InvitationRepository;
import com.brahimbh18.tictactoesvx.data.repository.MatchRepository;
import com.brahimbh18.tictactoesvx.data.repository.UserRepository;

public final class ServiceLocator {
    private static TokenStorage tokenStorage;
    private static ApiService apiService;
    private static AuthRepository authRepository;
    private static UserRepository userRepository;
    private static FriendRepository friendRepository;
    private static InvitationRepository invitationRepository;
    private static MatchRepository matchRepository;

    public static void init(Context context) {
        if (tokenStorage == null) {
            tokenStorage = new TokenStorage(context.getApplicationContext());
            apiService = new ApiClient();
            authRepository = new AuthRepository(tokenStorage);
            userRepository = new UserRepository(tokenStorage);
            friendRepository = new FriendRepository();
            invitationRepository = new InvitationRepository();
            matchRepository = new MatchRepository();
        }
    }

    public static TokenStorage tokenStorage() { return tokenStorage; }
    public static ApiService apiService() { return apiService; }
    public static AuthRepository authRepository() { return authRepository; }
    public static UserRepository userRepository() { return userRepository; }
    public static FriendRepository friendRepository() { return friendRepository; }
    public static InvitationRepository invitationRepository() { return invitationRepository; }
    public static MatchRepository matchRepository() { return matchRepository; }

    private ServiceLocator() {}
}
