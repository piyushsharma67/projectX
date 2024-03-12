import { createAsyncThunk } from "@reduxjs/toolkit";
import ServiceApi from "../../../../utils/httpService";
import { ApiEndpoints } from "../../../../utils/apiEndpoints";

export type SignUpBody = {
    name: string,
    email: string,
    password: string,
    mobile: string
}

export const signUpThunk = createAsyncThunk(
    "signUpThunk", async (body: SignUpBody, thunkAPI) => {
        console.log("i am called", body)
        try {
            const { status, data } = await ServiceApi.getInstance().POST(
                ApiEndpoints.authentication.signUp,
                body,
                {
                    headers: {
                    },
                }
            )
            if (status != 200 || !data.status) {
                throw new Error(data.message)
            }
            console.log("data is", data)
            return data
        } catch (error: any) {
            console.log("error is", error)
            return thunkAPI.rejectWithValue(error.message)
        }
    }
)

export type LoginBody = {
    email: string,
    password: string,
}

export const loginThunk = createAsyncThunk(
    "loginThunk", async (body: LoginBody, thunkAPI) => {
        try {

            const { status, data } = await ServiceApi.getInstance().POST(
                ApiEndpoints.authentication.login,
                body,
                {
                    headers: {
                        // Authorization: `Bearer ${token}`,
                    },
                }
            )
            if (status != 200 || !data.status) {
                throw new Error(data.message)
            }
            return true
        } catch (error: any) {
            return thunkAPI.rejectWithValue(error.message)
        }
    }
)