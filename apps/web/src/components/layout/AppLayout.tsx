import type { ReactNode } from 'react';

interface AppLayoutProps {
    children: ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {
    return (
        <div className="app-shell">
            <a className="usa-skipnav" href="#main-content">
                Skip to main content
            </a>

            <header className="usa-header usa-header--basic">
                <div className="usa-nav-container">
                    <div className="usa-navbar">
                        <div className="usa-logo" id="basic-logo">
                            <a href="/" title="Fitness Diet Tracker home">
                                Fitness Diet Tracker
                            </a>
                        </div>
                    </div>

                    <nav aria-label="Primary navigation" className="usa-nav">
                        <ul className="usa-nav__primary usa-accordion">
                            <li className="usa-nav__primary-item">
                                <a href="/" className="usa-nav-link">
                                    <span>Dashboard</span>
                                </a>
                            </li>
                            <li className="usa-nav__primary-item">
                                <a href="/exercises" className="usa-nav-link">
                                    <span>Exercises</span>
                                </a>
                            </li>
                            <li className="usa-nav__primary-item">
                                <a href="/meals" className="usa-nav-link">
                                    <span>Meals</span>
                                </a>
                            </li>
                            <li className="usa-nav__primary-item">
                                <a href="/goals" className="usa-nav-link">
                                    <span>Goals</span>
                                </a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </header>

            <main id="main-content" className="app-main grid-container">
                {children}
            </main>

            <footer className="usa-footer usa-footer--slim">
                <div className="grid-container usa-footer__return-to-top">
                    <a href="#main-content">Return to top</a>
                </div>
            </footer>
        </div>
    );
}
