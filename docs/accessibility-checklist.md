# Accessibility Checklist

## Overview

This checklist defines baseline accessibility expectations for the Fitness Diet Tracker frontend.

The goal is to keep the app usable for people relying on keyboards, screen readers, mobile devices, clear visual hierarchy, and accessible form interactions. The app should follow USWDS-inspired accessibility and usability practices as the UI develops.

## Page Structure

- [ ] Each page has one clear `<h1>`.
- [ ] Headings follow a logical order without skipping levels unnecessarily.
- [ ] Main content is wrapped in a `<main>` landmark.
- [ ] Navigation is wrapped in a `<nav>` landmark with an accessible label.
- [ ] Repeated layout elements such as header, navigation, main content, and footer are consistent across pages.
- [ ] Pages include a skip link that allows keyboard users to jump to main content.

## Keyboard Navigation

- [ ] All interactive elements can be reached with the keyboard.
- [ ] Keyboard focus order follows the visual order of the page.
- [ ] Buttons, links, form controls, and navigation items show visible focus states.
- [ ] No keyboard trap is introduced.
- [ ] Modals, menus, and dropdowns can be opened, used, and closed with the keyboard when they are added.

## Forms

- [ ] Every input has a visible label.
- [ ] Labels are programmatically associated with inputs using `htmlFor` and `id`.
- [ ] Required fields are clearly identified.
- [ ] Form instructions are placed before the relevant inputs.
- [ ] Error messages explain what went wrong and how to fix it.
- [ ] Error messages are associated with the related field when possible.
- [ ] Inputs use appropriate types such as `email`, `number`, `date`, or `text`.

## Buttons and Links

- [ ] Buttons are used for actions.
- [ ] Links are used for navigation.
- [ ] Link text clearly describes the destination.
- [ ] Button text clearly describes the action.
- [ ] Icon-only buttons include accessible labels.

## Color and Contrast

- [ ] Text has sufficient contrast against its background.
- [ ] Important information is not communicated by color alone.
- [ ] Error, warning, success, and status messages include text or icons in addition to color.
- [ ] Focus states are clearly visible.

## Responsive Design

- [ ] Layout works at mobile, tablet, and desktop widths.
- [ ] Content does not require horizontal scrolling on small screens.
- [ ] Buttons and form controls are large enough for touch interaction.
- [ ] Navigation remains usable on smaller screens.
- [ ] Cards, grids, and dashboard sections stack clearly on mobile.

## Content Clarity

- [ ] Page titles are clear and specific.
- [ ] Instructions are written in plain language.
- [ ] Labels use user-friendly terms.
- [ ] Error messages are specific and actionable.
- [ ] Fitness and nutrition terms are explained when needed.

## Dynamic Content

- [ ] Loading states are communicated clearly.
- [ ] Empty states explain what the user can do next.
- [ ] Success states confirm completed actions.
- [ ] Error states give recovery steps.
- [ ] Future charts or progress visuals include text summaries.

## Fitness and Nutrition Specific Checks

- [ ] Exercise forms clearly explain required fields.
- [ ] Meal forms clearly label calories and macro fields.
- [ ] Units are clearly shown, such as minutes, pounds, grams, or calories.
- [ ] Goal progress is described in text, not only shown visually.
- [ ] Dietary restrictions are displayed clearly and respectfully.
- [ ] Nutrition guidance avoids presenting itself as medical advice.

## Pull Request Review Checklist

Before merging frontend UI changes, confirm:

- [ ] The page or component works with keyboard navigation.
- [ ] Form fields have labels.
- [ ] Buttons and links use the correct semantic element.
- [ ] The layout works on mobile and desktop.
- [ ] Main content is easy to identify.
- [ ] Error and empty states are considered.
- [ ] New UI follows the project’s USWDS-inspired design direction.

## Notes

This checklist should be updated as the app gains more complex interactions, including authentication, dashboards, charts, modals, meal logging, exercise logging, and goal tracking.
