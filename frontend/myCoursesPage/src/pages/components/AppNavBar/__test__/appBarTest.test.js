import React from 'react'
import { Provider } from 'react-redux'
import store, { setupStore } from '../../../../redux/store'
import AppNavBarIndex from '../index'
import { screen, render } from "@testing-library/react"
import { MemoryRouter as Router } from 'react-router-dom'
import '@testing-library/jest-dom'

describe("test the appbar component", () => {

    it('if user is not available then login button is visible', () => {
        render(
            <Router initialEntries={['/']} initialIndex={0}>
                <Provider store={store}>
                    <AppNavBarIndex />
                </Provider>
            </Router>
        )
        const linkElement = screen.getByText("Login");
        expect(linkElement).toBeInTheDocument();
    })

    it('if user is available then Avatar is visible', () => {
        const preloadedState = {
            authReducer: {
                user: {
                    name: "piyush",
                    email: "test@g.co",
                    mobile: "32132"
                }
            }
        }
        let store = setupStore(preloadedState)
        render(
            <Router initialEntries={['/']} initialIndex={0}>
                <Provider store={store}>
                    <AppNavBarIndex />
                </Provider>
            </Router>
        )
        const linkElement = screen.getByTestId("avatar");
        expect(linkElement).toBeInTheDocument();
    })
})