const { renderHook } = require("@testing-library/react");
const { describe } = require("node:test");
import useAuthentication from '../hooks/useAuthentication'
import UserAuthenticationPage from '../userAuthenticationPage'
import { render, fireEvent, screen, waitFor, act } from '@testing-library/react';
import store from '../../../redux/store'
import { MemoryRouter as Router } from 'react-router-dom'
import '@testing-library/jest-dom'
import { Provider } from 'react-redux';
import React from 'react'

describe('test the useAuthenticationHook', () => {
    it('test the onSelected function when login button is clicked ', () => {
        const { result } = renderHook(() => useAuthentication(), {
            wrapper: ({ children }) => (
                <Router initialEntries={['/']} initialIndex={0}>
                    <Provider store={store}>{children}</Provider>
                </Router>
            ),
        })
        render(
            <Router initialEntries={['/']} initialIndex={0}>
                <Provider store={store}>
                    <UserAuthenticationPage forms={['Login', 'Signup']} selectedFormIndex={2} onSelectedForm={result.current.onSelectedForm} />
                </Provider>
            </Router>
        )
        const button = screen.getByTestId('Login-0')
        fireEvent.click(button)
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
})