import React from "react";
import styles from './footer.module.css'
import logo from '../../../assets/logo.png'

function Footer() {
    return (
        <div className={styles.footerContainer}>
            <div className={styles.block1}>
                <div className={styles.logoContainer}>
                    <img src={logo} />
                </div>
                <div className={styles.sloganContainer}>
                    <span className={styles.sloganText}>Virtual class for zoom</span>
                </div>
            </div>
            <div className={styles.block2}>
                <div className={styles.block2_inner1}>
                    <span>Careers</span>
                    <span>Privacy policy</span>
                    <span>terms and condition</span>
                </div>
                <div className={styles.block2_inner2}>
                    <span>© 2021 Class Technologies Inc. </span>
                </div>
            </div>
        </div>
    )
}

export default React.memo(Footer)