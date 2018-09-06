import * as React from "react"
import axios from "axios";
import * as cookie from "js-cookie";

// components
import Head from "./Head";
import Foot from "./Foot";
import PlayList from "./PlayList";
import Loading from "./Loading";
import Map from "./Map";

interface State {
  loading: boolean;
  cityName: string;
  weather: string;
  playlists;
  latitude: number;
  longitude: number;
}

export default class App extends React.Component<any, State> {

  componentDidMount() {
    this.fetchTracks()
  }

  constructor(props) {
    super(props);

    this.state = {
      loading: true,
      cityName: "",
      weather: "",
      playlists: [],
      latitude: 59.95,
      longitude: 30.33
    }
  }

  render() {
    const comp = this.state.loading ? <Loading/> : (
      <>
        <PlayList weather={this.state.weather} cityName={this.state.cityName} playlists={this.state.playlists} />
        <Map lat={this.state.latitude} lng={this.state.longitude}/>
      </>
    );

    return (
      <>
        <Head/>
          {comp}
        <Foot/>
      </>
    )
  }

  fetchTracks = () => {
    navigator.geolocation.getCurrentPosition(
      async data => {
        const { latitude, longitude } = data.coords;
        const res = await axios.get("/api/tracks", {
          params: {
            latitude,
            longitude
          }
        });
        this.setState({
          loading: false,
          cityName: res.data.city_name,
          weather: res.data.weather,
          playlists: res.data.playlist.playlists.items || [],
          latitude,
          longitude
        })
      },
      err => console.error(err)
    )
  }

  logout() {
    cookie.remove("musi-token");
    location.reload();
  }

}