import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import api from "../services/api";

export default function EntitiesPage() {

  const [entities, setEntities] =
    useState([]);

  useEffect(() => {
    loadEntities();
  }, []);

  async function loadEntities() {

    const res =
      await api.get("/entities");

    //    console.log("ENTITIES RESPONSE", res.data)

    setEntities(res.data);
  }

  return (
    // <div>

    //   <h1>Entities</h1>

    //   <Link to="/entities/new">
    //     Create Entity
    //   </Link>

    //   <hr />

    //   {entities.map(entity => (

    //     <div key={entity.id}>

    //       <h3>{entity.name}</h3>

    //       <Link
    //         to={`/entities/${entity.id}/news`}
    //       >
    //         View News
    //       </Link>

    //     </div>
    //   ))}
    // </div>

    // temp 
    <div>

    <h1>Entities</h1>

    <Link to="/entities/new">
      Create Entity
    </Link>

    <hr />

    {entities.map(entity => (
      <div key={entity.ID}>

        <h3>{entity.Name}</h3>

        <p>ID: {entity.id}</p>

        <Link
          to={`/entities/${entity.ID}/news`}
        >
          View News
        </Link>

      </div>
    ))}

  </div>
  );
}