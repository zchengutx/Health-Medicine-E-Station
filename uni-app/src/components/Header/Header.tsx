import React from 'react'
import styles from './Header.module.css'

export interface HeaderProps {
  title: string
  onBack: () => void
  showMenuButton?: boolean
}

export const Header: React.FC<HeaderProps> = ({
  title,
  onBack,
  showMenuButton = false
}) => {
  return (
    <header className={styles.header}>
      <button
        className={styles.backButton}
        onClick={onBack}
        aria-label="返回"
      >
        <svg
          className={styles.backIcon}
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        >
          <path d="M19 12H5" />
          <path d="M12 19l-7-7 7-7" />
        </svg>
      </button>
      
      <h1 className={styles.title}>{title}</h1>
      
      {showMenuButton && (
        <button
          className={styles.menuButton}
          aria-label="菜单"
        >
          <svg
            className={styles.menuIcon}
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <circle cx="12" cy="12" r="1" />
            <circle cx="19" cy="12" r="1" />
            <circle cx="5" cy="12" r="1" />
          </svg>
        </button>
      )}
    </header>
  )
}

export default Header