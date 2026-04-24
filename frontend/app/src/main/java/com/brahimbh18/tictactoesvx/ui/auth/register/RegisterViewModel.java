package com.brahimbh18.tictactoesvx.ui.auth.register;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;
import com.brahimbh18.tictactoesvx.core.util.Result;
import com.brahimbh18.tictactoesvx.ui.common.UiState;

public class RegisterViewModel extends ViewModel {
    private final MutableLiveData<UiState<String>> registerState = new MutableLiveData<>(UiState.idle());

    public LiveData<UiState<String>> getRegisterState() {
        return registerState;
    }

    public void register(String username, String password) {
        registerState.setValue(UiState.loading());
        Result<String> result = ServiceLocator.authRepository().register(username, password);
        if (result.isSuccess()) {
            registerState.setValue(UiState.success(result.getData()));
        } else {
            registerState.setValue(UiState.error(result.getError()));
        }
    }
}
