import React, {useEffect} from 'react';

import {Route, withRouter, Redirect, Switch, Link} from 'react-router-dom';




import {Loading, IFrame} from "../../styles/project";
import {classes} from "../../utils/Router";
import {client} from "../../utils/requests";
import connect from "react-redux/es/connect/connect";

import CircularProgress from "@material-ui/core/CircularProgress";
import Divider from "@material-ui/core/Divider";
import moment from "moment";
import Tooltip from "@material-ui/core/Tooltip";

moment.locale('zh-CN');



const Project = (props) => {
  const dispatch = (type, payload) => props.dispatch({type, payload});

  const breadcrumb = [{label:'/project', value: '项目'}];
  
  dispatch('breadcrumb', breadcrumb);
  const {url} = props.match.params;


  const [data, setData] = React.useState({});
  const [loading, setLoading] = React.useState(false);
  
  useEffect(()=>{
    setLoading(true)
    client.get(`/Project/${url}`)
      .then(r => {
          if (r.data.ret) {
            const data = r.data.res;
            setData(data)
            dispatch('breadcrumb', breadcrumb.concat([{label: props.match.url, value: data.Name}]));
          }
          
        }
      ).catch(e => {
      console.log(e);
    }).finally(()=> {
      setLoading(false)
    })
    
  }, [url]);
  


  return (
        <>
          {
            loading && <Loading><CircularProgress  /></Loading>
          }

            {
              data.hasOwnProperty('Frame') &&
                <IFrame frameBorder='no'  src={data.Frame} />

            }
  

          
        </>
    )

}

const mapStateToProps = state => ({

});
export default connect(
  mapStateToProps,
)(withRouter(Project));
