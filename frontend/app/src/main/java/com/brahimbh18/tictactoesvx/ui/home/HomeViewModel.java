package com.brahimbh18.tictactoesvx.ui.home;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.domain.model.User;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

public class HomeViewModel extends ViewModel {
    private final MutableLiveData<UiState<User>> meState = new MutableLiveData<>(UiState.idle());

    public LiveData<UiState<User>> getMeState() {
        return meState;
    }

    public void loadMe() {
        meState.setValue(UiState.loading());
        Result<User> result = ServiceLocator.userRepository().getCurrentUser();
        if (result.isSuccess()) {
            meState.setValue(UiState.success(result.getData()));
        } else {
            meState.setValue(UiState.error(result.getError()));
        }
    }

    public void logout() {
        ServiceLocator.authRepository().logout();
    }
}
