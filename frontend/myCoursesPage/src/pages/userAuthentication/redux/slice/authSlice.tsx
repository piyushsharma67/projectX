import { createSlice } from '@reduxjs/toolkit';
import { initialStateAuth } from '../types';
import { signUpThunk } from '../thunks/authentication';

// Initial state
const initialState = {
    user: {
        email: "",
        name: "",
        mobile: "",
    },
    token: "",
    isSigningUp: false
} as initialStateAuth

// Create a slice
const authSlice = createSlice({
    name: 'counter',
    initialState,
    extraReducers(builder) {
        builder.addCase(signUpThunk.pending, (state) => {
            state.isSigningUp = true
        })
        builder.addCase(signUpThunk.rejected, (state) => {
            state.isSigningUp = false
        })
        builder.addCase(signUpThunk.fulfilled, (state, action) => {
            console.log("payload is", action.payload)
            state.user = { ...action.payload.data.User }
            state.token = action.payload.data.Token
        })
    },
    reducers: {

    },
});

// Export the generated action creators
// export const { increment, decrement, reset, } = authSlice.actions;

// Export the reducer
export default authSlice.reducer;