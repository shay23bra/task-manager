import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom"; // Import the custom matchers
import App from "./App";

test("renders Task Manager heading", () => {
  render(<App />);
  const headingElement = screen.getByText(/Task Manager/i);
  expect(headingElement).toBeInTheDocument();
});
