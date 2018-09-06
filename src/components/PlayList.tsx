import * as React from "react"
import PlayListItem from "./PlayListItem";
import ReloadView from "./Reload";
import styled from "styled-components";

const CityName = styled.h1`
  font-size: 2rem;
  letter-spacing: 2px;
`

const Weather = styled.p`
  font-size: 1.5rem;
  letter-spacing: 2px;
`

const List = styled.ul`
  list-style: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  padding-bottom: 10px;
`

const Flex = styled.div`
  margin: 10px;
  display: flex;
  flex-direction: column;
`

interface Props {
  cityName: string;
  weather: string;
  playlists: any[];
}

export default class PlayList extends React.Component<Props> {
  render() {
    return(
      <>
        {
          this.props.playlists.length === 0 ? (
            <ReloadView/>
          ):(
            <>
              <Flex>
                <CityName>{this.props.cityName.toUpperCase()}</CityName>
                <Weather>{this.props.weather}</Weather>
              </Flex>
              <List>
                {this.props.playlists.map(item => <PlayListItem item={item} key={item.id}/>)}
              </List>
            </>
          )
        }
      </>
    )
  }
}