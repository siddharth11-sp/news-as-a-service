import {
  useEffect,
  useState,
} from "react";

import {
  useParams,
} from "react-router-dom";

import api from "../services/api";

export default function NewsFeedPage() {

  const { id } = useParams();

  const [news, setNews] =
    useState([]);

  async function loadNews() {

    const res =
      await api.get(
        `/entities/${id}/news`
      );

    setNews(res.data.data);
  }

  async function refresh() {

    await api.post(
      `/entities/${id}/refresh`
    );

    alert(
      "Refresh started"
    );
  }

  useEffect(() => {
    loadNews();
  }, []);

  return (
    <div>

      <h1>News Feed</h1>

      <button onClick={refresh}>
        Refresh
      </button>

      <br />
      <br />

      <table border="1">

        <thead>

          <tr>
            <th>Title</th>
            <th>Source</th>
            <th>Date</th>
            <th>Sentiment</th>
          </tr>

        </thead>

        <tbody>

          {news.map(item => (

            <tr key={item.id}>

              <td>
                <a
                  href={item.url}
                  target="_blank"
                >
                  {item.Title}
                </a>
              </td>

              <td>
                {item.Source}
              </td>

              <td>
                {item.PublishedDate}
              </td>

              <td>
                {item.Sentiment}
              </td>

            </tr>

          ))}

        </tbody>

      </table>

    </div>
  );
}