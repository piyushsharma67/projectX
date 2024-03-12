import React from 'react'
import useAuthentication from './hooks/useAuthentication'
import SignUpPageIndex from './view'

function UserAuthenticationPage() {
    const {
        token,
        loading,
        forms,
        errors,
        selectedFormIndex,
        register,
        onSelectedForm,
        onClickLogin,
        onClickSignUp
    } = useAuthentication()

    return (
        <SignUpPageIndex
            token={token}
            loading={loading}
            forms={forms}
            selectedFormIndex={selectedFormIndex}
            errors={errors}
            onSelectedForm={onSelectedForm}
            register={register}
            onClickLogin={onClickLogin}
            onClickSignUp={onClickSignUp}
        />
    )
}

export default UserAuthenticationPage