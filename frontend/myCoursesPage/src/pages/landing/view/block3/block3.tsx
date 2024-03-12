import InfoCard from '../../../components/infoCard/InfoCard'
import styles from './block3.module.css'
import ArticleIcon from '@mui/icons-material/Article';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import PeopleIcon from '@mui/icons-material/People';

function Block3() {
    return (
        <div className={styles.block3Container}>
            <div className={styles.headingContainer}>
                <p className={styles.heading}>All-In-One <span>Cloud Software.</span></p>
                <span className={styles.subHeading}>TOTC is one powerful online software suite that combines all the tools needed to run a successful school or office.</span>
            </div>

            <div className={styles.cardContainer}>
                <InfoCard
                    headingText='Online Billing, Invoicing, & Contracts'
                    subHeadingText='Simple and secure control of your organization’s financial and legal transactions. Send customized invoices and contracts'
                >
                    <div className={styles.iconContainer}>
                        <ArticleIcon color='info' />
                    </div>
                </InfoCard>
                <InfoCard
                    headingText='Online Billing, Invoicing, & Contracts'
                    subHeadingText='Schedule and reserve classrooms at one campus or multiple campuses. Keep detailed records of student attendance'
                >
                    <div className={styles.iconContainer} style={{ backgroundColor: '#00CBB8' }}>
                        <CalendarMonthIcon color='info' />
                    </div>
                </InfoCard>
                <InfoCard
                    headingText='Customer Tracking'
                    subHeadingText='Automate and track emails to individuals or groups. Skilline’s built-in system helps organize your organization '
                >
                    <div className={styles.iconContainer} style={{ backgroundColor: '#29B9E7' }}>
                        <PeopleIcon color='info' />
                    </div>
                </InfoCard>
            </div>
            <div className={styles.textImageContainer}>

            </div>
        </div>
    )
}

export default Block3