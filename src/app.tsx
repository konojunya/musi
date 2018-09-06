import * as React from "react"
import { hydrate } from "react-dom";

// app
import App from "./components/App";

hydrate(
  <App/>,
  document.getElementById("app")
)