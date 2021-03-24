/*!

=========================================================
* Light Bootstrap Dashboard React - v1.3.0
=========================================================

* Product Page: https://www.creative-tim.com/product/light-bootstrap-dashboard-react
* Copyright 2019 Creative Tim (https://www.creative-tim.com)
* Licensed under MIT (https://github.com/creativetimofficial/light-bootstrap-dashboard-react/blob/master/LICENSE.md)

* Coded by Creative Tim

=========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

*/
import Dashboard from "views/Dashboard.jsx";
import IngestionTool from "views/IngestionTool.jsx";
import ToDoList from "views/ToDoList.jsx";
import ClaimsTool from "views/ClaimsTool.jsx";
import OwnershipConflict from "views/OwnershipConflict.jsx";
import Maps from "views/Maps.jsx";
import Analytics from "views/Analytics.jsx";
import Reporting from "views/Reporting.jsx";

const dashboardRoutes = [
  {
    path: "/dashboard",
    name: "Home",
    icon: "pe-7s-music",
    component: Dashboard,
    layout: "/admin"
  },
  {
    path: "/todo",
    name: "To Do List",
    icon: "pe-7s-timer",
    component: ToDoList,
    layout: "/admin"
  },
  {
    path: "/claims",
    name: "Claims Tool",
    icon: "pe-7s-cash",
    component: ClaimsTool,
    layout: "/admin"
  },
  {
    path: "/ingestion",
    name: "Ingestion Tool",
    icon: "pe-7s-news-paper",
    component: IngestionTool,
    layout: "/admin"
  },
  {
    path: "/ownership",
    name: "Ownership Conflict",
    icon: "pe-7s-users",
    component: OwnershipConflict,
    layout: "/admin"
  },
  {
    path: "/analytics",
    name: "Analytics",
    icon: "pe-7s-graph1",
    component: Analytics,
    layout: "/admin"
  },
  {
    path: "/reporting",
    name: "Reporting",
    icon: "pe-7s-server",
    component: Reporting,
    layout: "/admin"
  }
  /*
  {
    path: "/maps",
    name: "Maps",
    icon: "pe-7s-map-marker",
    component: Maps,
    layout: "/admin"
  },
  */
];

export default dashboardRoutes;
