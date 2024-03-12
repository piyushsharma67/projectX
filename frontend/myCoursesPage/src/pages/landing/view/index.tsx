import Block2 from "./block2/block2"
import Block1 from "./block1/block1"
import Block3 from "./block3/block3"
import Block4 from "./block4/block4"
import Footer from "../../components/footer/footer"

function LandingPageIndex() {
    return (
        <div style={{ width: '100%' }}>
            <Block1 />
            <Block2 />
            <Block3 />
            <Block4 />
            <Footer />
        </div>
    )
}

export default LandingPageIndex