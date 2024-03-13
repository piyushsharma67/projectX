const { renderHook } = require("@testing-library/react");
const { describe } = require("node:test");
import useAuthentication from '../hooks/useAuthentication'
import { waitFor } from '@testing-library/react';
import store from '../../../redux/store'
import { MemoryRouter as Router } from 'react-router-dom'
import '@testing-library/jest-dom'
import { Provider } from 'react-redux';
import React from 'react'
import { signUpThunk } from '../redux/thunks/authentication'

jest.mock("../redux/thunks/authentication", () => ({
    signUpThunk: jest.fn()
})
)

describe('test the useAuthenticationHook', () => {
    it('test the onSelected function when login button is clicked ', async () => {
        const { result } = renderHook(() => useAuthentication(), {
            wrapper: ({ children }) => (
                <Router initialEntries={['/']} initialIndex={0}>
                    <Provider store={store}>{children}</Provider>
                </Router>
            ),
        })

        await waitFor(() => {
            result.current.onSelectedForm(1)()
        })

        expect(result.current.selectedFormIndex).toBe(1)
    })

    it('test the onSelected function when signup button is selected', async () => {
        const { result } = renderHook(() => useAuthentication(), {
            wrapper: ({ children }) => (
                <Router initialEntries={['/']} initialIndex={0}>
                    <Provider store={store}>{children}</Provider>
                </Router>
            ),
        })

        await waitFor(() => {
            result.current.onSelectedForm(2)()
        })

        await waitFor(async () => {
            expect(result.current.selectedFormIndex).toBe(2)
        })
    })

    it('test the onClickSignUp function when signup button is selected', async () => {

        render()
        const { result } = renderHook(() => useAuthentication(), {
            wrapper: ({ children }) => (
                <Router initialEntries={['/']} initialIndex={0}>
                    <Provider store={store}>{children}</Provider>
                </Router>
            ),
        })

        signUpThunk.mockResolvedValue({ meta: { requestStatus: 'fulfilled' } });


    })
})