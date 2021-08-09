import React from "react";

import { PublishedAppProvider as ThemeProvider } from "./theme";
import { Provider as LocalStateProvider } from "./state";
import { PublishedProvider as IntlProvider } from "./locale";
import { Provider as DndProvider } from "./util/use-dnd";

import PublishedPage from "@reearth/components/pages/Published";

export default function App() {
  return (
    <ThemeProvider>
      <LocalStateProvider>
        <DndProvider>
          <IntlProvider>
            <PublishedPage />
          </IntlProvider>
        </DndProvider>
      </LocalStateProvider>
    </ThemeProvider>
  );
}
