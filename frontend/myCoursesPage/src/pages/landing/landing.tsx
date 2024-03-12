import style from './landing.module.css'
import LandingPageIndex from "./view"

function LandingPage() {
    return (
        <div className={style.page}>
            <LandingPageIndex />
        </div>
    )
}

export default LandingPage