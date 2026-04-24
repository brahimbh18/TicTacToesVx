package com.brahimbh18.tictactoesvx.ui.auth.login;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

public class LoginViewModel extends ViewModel {
    private final MutableLiveData<UiState<String>> loginState = new MutableLiveData<>(UiState.idle());

    public LiveData<UiState<String>> getLoginState() {
        return loginState;
    }

    public void login(String username, String password) {
        loginState.setValue(UiState.loading());
        Result<String> result = ServiceLocator.authRepository().login(username, password);
        if (result.isSuccess()) {
            loginState.setValue(UiState.success(result.getData()));
        } else {
            loginState.setValue(UiState.error(result.getError()));
        }
    }
}
