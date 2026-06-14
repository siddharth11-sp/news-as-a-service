import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

import EntitiesPage from "./pages/EntitiesPage";
import CreateEntityPage from "./pages/CreateEntityPage";
import NewsFeedPage from "./pages/NewsFeedPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>

        <Route
          path="/"
          element={<EntitiesPage />}
        />

        <Route
          path="/entities/new"
          element={<CreateEntityPage />}
        />

        <Route
          path="/entities/:id/news"
          element={<NewsFeedPage />}
        />

      </Routes>
    </BrowserRouter>
  );
}

export default App;