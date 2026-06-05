export function DashboardPage() {
    return (
        <section className="usa-section">
            <div className="grid-row grid-gap">
                <div className="tablet:grid-col-8">
                    <h1>Fitness and nutrition dashboard</h1>
                    <p className="usa-intro">
                        Track exercise, meals, nutrition goals, and dietary restrictions in
                        one place.
                    </p>
                </div>
            </div>

            <ul className="usa-card-group margin-top-4">
                <li className="usa-card tablet:grid-col-4">
                    <div className="usa-card__container">
                        <div className="usa-card__header">
                            <h2 className="usa-card__heading">Today&apos;s exercise</h2>
                        </div>
                        <div className="usa-card__body">
                            <p>Log workouts, duration, sets, reps, and notes.</p>
                        </div>
                        <div className="usa-card__footer">
                            <button type="button" className="usa-button">
                                Add exercise
                            </button>
                        </div>
                    </div>
                </li>

                <li className="usa-card tablet:grid-col-4">
                    <div className="usa-card__container">
                        <div className="usa-card__header">
                            <h2 className="usa-card__heading">Today&apos;s meals</h2>
                        </div>
                        <div className="usa-card__body">
                            <p>Track calories, protein, carbs, fat, and meal notes.</p>
                        </div>
                        <div className="usa-card__footer">
                            <button type="button" className="usa-button">
                                Add meal
                            </button>
                        </div>
                    </div>
                </li>

                <li className="usa-card tablet:grid-col-4">
                    <div className="usa-card__container">
                        <div className="usa-card__header">
                            <h2 className="usa-card__heading">Goal progress</h2>
                        </div>
                        <div className="usa-card__body">
                            <p>Review progress toward your fitness and nutrition goals.</p>
                        </div>
                        <div className="usa-card__footer">
                            <button type="button" className="usa-button usa-button--outline">
                                View goals
                            </button>
                        </div>
                    </div>
                </li>
            </ul>
        </section>
    );
}
