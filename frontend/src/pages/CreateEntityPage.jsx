import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

export default function CreateEntityPage() {

  const navigate = useNavigate();

  const [name, setName] =
    useState("");

  const [aliases, setAliases] =
    useState("");

  const [interval, setInterval] =
    useState(15);

  async function submit(e) {

    e.preventDefault();

    await api.post(
      "/entities",
      {
        name,
        aliases:
          aliases
            .split(",")
            .map(x => x.trim())
            .filter(Boolean),
        ingestion_interval_minutes:
          Number(interval),
      }
    );

    navigate("/");
  }

  return (
    <div>

      <h1>Create Entity</h1>

      <form onSubmit={submit}>

        <input
          placeholder="Name"
          value={name}
          onChange={(e) =>
            setName(e.target.value)
          }
        />

        <br />

        <input
          placeholder="Aliases"
          value={aliases}
          onChange={(e) =>
            setAliases(e.target.value)
          }
        />

        <br />

        <input
          type="number"
          value={interval}
          onChange={(e) =>
            setInterval(e.target.value)
          }
        />

        <br />

        <button type="submit">
          Create
        </button>

      </form>

    </div>
  );
}