package com.brahimbh18.tictactoesvx.core.network;

import com.brahimbh18.tictactoesvx.core.util.Constants;

public class ApiClient implements ApiService {
    @Override
    public String getBaseUrl() {
        return Constants.BASE_URL;
    }
}
