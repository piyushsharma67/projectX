import { Action, ThunkDispatch, combineReducers, configureStore } from "@reduxjs/toolkit";
import { useDispatch } from "react-redux";
import authSlice from "../pages/userAuthentication/redux/slice/authSlice";

const rootReducers = combineReducers({
    authReducer: authSlice
})
const store = configureStore({
    reducer: rootReducers
})

export default store

export type RootAppState = ReturnType<typeof store.getState>
export type AppDispatch = ThunkDispatch<RootAppState, any, Action>

export const useAppDispatch = () => useDispatch<AppDispatch>()

export function setupStore(preloadedState?: Partial<RootAppState>) {
    return configureStore({
        reducer: rootReducers,
        preloadedState
    })
}