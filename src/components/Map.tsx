import * as React from "react"
import GoogleMapReact from 'google-map-react';
import styled from "styled-components"

const Wrapper = styled.div`
  width: 100%;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;
  z-index: -100;
  opacity: 0.4;
`

interface Props {
  lat;
  lng;
  zoom;
}

export default class Map extends React.Component<Props> {
  static defaultProps = {
    zoom: 13
  }

  render() {
    return (
      <Wrapper>
        <GoogleMapReact
          bootstrapURLKeys={{ key: "" }}
          defaultCenter={{ lat: this.props.lat, lng: this.props.lng }}
          defaultZoom={this.props.zoom}
        />
      </Wrapper>
    )
  }
}