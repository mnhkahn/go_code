


<!DOCTYPE html>
<html>
  <head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# githubog: http://ogp.me/ns/fb/githubog#">
    <meta charset='utf-8'>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>gobook/dive-into/interface/02/interface.c at master · xushiwei/gobook</title>
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub" />
    <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub" />
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />

    
    

    <meta content="authenticity_token" name="csrf-param" />
<meta content="TKu//ElHDFUrBXmxgvW0PEc9tjeWwIdFMXgxypzof3U=" name="csrf-token" />

    <link href="https://a248.e.akamai.net/assets.github.com/stylesheets/bundles/github-9b7be9f578b4af170903c220215ad8680a97d61e.css" media="screen" rel="stylesheet" type="text/css" />
    <link href="https://a248.e.akamai.net/assets.github.com/stylesheets/bundles/github2-02faadd95ff1183c3df5da2e69d237d00cef71b0.css" media="screen" rel="stylesheet" type="text/css" />
    

    <script src="https://a248.e.akamai.net/assets.github.com/javascripts/bundles/frameworks-7b5694dece50ddf8456fccf7884bd83581722a3f.js" type="text/javascript"></script>
    
    <script defer="defer" src="https://a248.e.akamai.net/assets.github.com/javascripts/bundles/github-ed88cce616e7fb55dea93e5ffadc963054853f83.js" type="text/javascript"></script>
    

      <link rel='permalink' href='/xushiwei/gobook/blob/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into/interface/02/interface.c'>
    <meta property="og:title" content="gobook"/>
    <meta property="og:type" content="githubog:gitrepository"/>
    <meta property="og:url" content="https://github.com/xushiwei/gobook"/>
    <meta property="og:image" content="https://a248.e.akamai.net/assets.github.com/images/gravatars/gravatar-140.png?1329275975"/>
    <meta property="og:site_name" content="GitHub"/>
    <meta property="og:description" content="gobook"/>

    <meta name="description" content="gobook" />
  <link href="https://github.com/xushiwei/gobook/commits/master.atom" rel="alternate" title="Recent Commits to gobook:master" type="application/atom+xml" />

  </head>


  <body class="logged_in page-blob  vis-public env-production " data-blob-contribs-enabled="yes">
    <div id="wrapper">

    
    
    

      <div id="header" class="true clearfix">
        <div class="container clearfix">
          <a class="site-logo" href="https://github.com/organizations/qbox">
            <!--[if IE]>
            <img alt="GitHub" class="github-logo" src="https://a248.e.akamai.net/assets.github.com/images/modules/header/logov7.png?1323882799" />
            <img alt="GitHub" class="github-logo-hover" src="https://a248.e.akamai.net/assets.github.com/images/modules/header/logov7-hover.png?1324325436" />
            <![endif]-->
            <img alt="GitHub" class="github-logo-4x" height="30" src="https://a248.e.akamai.net/assets.github.com/images/modules/header/logov7@4x.png?1323882799" />
            <img alt="GitHub" class="github-logo-4x-hover" height="30" src="https://a248.e.akamai.net/assets.github.com/images/modules/header/logov7@4x-hover.png?1324325436" />
          </a>

              
    <div class="topsearch ">
<form accept-charset="UTF-8" action="/search" id="top_search_form" method="get"><input name="utf8" type="hidden" value="&#x2713;" />        <a href="/search" class="advanced-search tooltipped downwards" title="Advanced Search"><span class="mini-icon advanced-search"></span></a>
        <div class="search placeholder-field js-placeholder-field">
          <input type="text" class="search my_repos_autocompleter" id="global-search-field" name="q" results="5" spellcheck="false" autocomplete="off" data-autocomplete="my-repos-autocomplete" placeholder="Search&hellip;" />
          <div id="my-repos-autocomplete" class="autocomplete-results">
            <ul class="js-navigation-container"></ul>
          </div>
          <input type="submit" value="Search" class="button">
          <span class="mini-icon search-input"></span>
        </div>
        <input type="hidden" name="type" value="Everything" />
        <input type="hidden" name="repo" value="" />
        <input type="hidden" name="langOverride" value="" />
        <input type="hidden" name="start_value" value="1" />
</form>      <ul class="top-nav">
          <li class="explore"><a href="https://github.com/explore">Explore</a></li>
          <li><a href="https://gist.github.com">Gist</a></li>
          <li><a href="/blog">Blog</a></li>
        <li><a href="http://help.github.com">Help</a></li>
      </ul>
    </div>


            


  <div id="userbox">
    <div id="user">
      <a href="https://github.com/hughlv"><img height="20" src="https://secure.gravatar.com/avatar/1ab98317c89c6c5270e069b1caaf259c?s=140&amp;d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png" width="20" /></a>
      <a href="https://github.com/hughlv" class="name">hughlv</a>
    </div>
    <ul id="user-links">
      <li>
        <a href="/new" id="new_repo" class="tooltipped downwards" title="Create a New Repo"><span class="mini-icon new-repo"></span></a>
      </li>
      <li>
        <a href="/inbox/notifications" id="notifications" class="tooltipped downwards" title="Notifications">
          <span class="mini-icon notifications"></span>
          <span class="unread_count">∞</span>
        </a>
      </li>
      <li><a href="/settings/profile" id="settings" class="tooltipped downwards" title="Account Settings"><span class="mini-icon account-settings"></span></a></li>
      <li>
          <a href="/logout" data-method="post" id="logout" class="tooltipped downwards" title="Log Out"><span class="mini-icon logout"></span></a>
      </li>
    </ul>
  </div>



          
        </div>
      </div>

      

            <div class="site hfeed" itemscope itemtype="http://schema.org/WebPage">
      <div class="container hentry">
        <div class="pagehead repohead instapaper_ignore readability-menu">
        <div class="title-actions-bar">
          



              <ul class="pagehead-actions">


          <li class="js-toggler-container watch-button-container ">
            <a href="/xushiwei/gobook/toggle_watch" class="minibutton btn-watch watch-button js-toggler-target" data-method="post" data-remote="true" rel="nofollow"><span><span class="icon"></span>Watch</span></a>
            <a href="/xushiwei/gobook/toggle_watch" class="minibutton btn-watch unwatch-button js-toggler-target" data-method="post" data-remote="true" rel="nofollow"><span><span class="icon"></span>Unwatch</span></a>
          </li>

              <li><a href="/xushiwei/gobook/fork_select" class="minibutton btn-fork " rel="facebox nofollow"><span><span class="icon"></span>Fork</span></a></li>

              <li class="nspr">
                <a href="/xushiwei/gobook/pull/new/master" class="minibutton btn-pull-request "><span><span class="icon"></span>Pull Request</span></a>
              </li>


      <li class="repostats">
        <ul class="repo-stats">
          <li class="watchers ">
            <a href="/xushiwei/gobook/watchers" title="Watchers" class="tooltipped downwards">
              3
            </a>
          </li>
          <li class="forks">
            <a href="/xushiwei/gobook/network" title="Forks" class="tooltipped downwards">
              1
            </a>
          </li>
        </ul>
      </li>
    </ul>

          <h1 itemscope itemtype="http://data-vocabulary.org/Breadcrumb" class="entry-title">
            <span class="mini-icon public-repo"></span>
            <span class="author vcard">
<a href="/xushiwei" class="url fn" itemprop="url" rel="author">              <span itemprop="title">xushiwei</span>
              </a></span> /
            <strong><a href="/xushiwei/gobook" class="js-current-repository">gobook</a></strong>
          </h1>
        </div>

          

  <ul class="tabs">
    <li><a href="/xushiwei/gobook" class="selected" highlight="repo_sourcerepo_downloadsrepo_commitsrepo_tagsrepo_branches">Code</a></li>
    <li><a href="/xushiwei/gobook/network" highlight="repo_network">Network</a>
    <li><a href="/xushiwei/gobook/pulls" highlight="repo_pulls">Pull Requests <span class='counter'>0</span></a></li>

      <li><a href="/xushiwei/gobook/issues" highlight="repo_issues">Issues <span class='counter'>0</span></a></li>

      <li><a href="/xushiwei/gobook/wiki" highlight="repo_wiki">Wiki <span class='counter'>0</span></a></li>

    <li><a href="/xushiwei/gobook/graphs" highlight="repo_graphsrepo_contributors">Graphs</a></li>

  </ul>
 
<div class="frame frame-center tree-finder" style="display:none"
      data-tree-list-url="/xushiwei/gobook/tree-list/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee"
      data-blob-url-prefix="/xushiwei/gobook/blob/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee"
    >

  <div class="breadcrumb">
    <span class="bold"><a href="/xushiwei/gobook">gobook</a></span> /
    <input class="tree-finder-input js-navigation-enable" type="text" name="query" autocomplete="off" spellcheck="false">
  </div>

    <div class="octotip">
      <p>
        <a href="/xushiwei/gobook/dismiss-tree-finder-help" class="dismiss js-dismiss-tree-list-help" title="Hide this notice forever" rel="nofollow">Dismiss</a>
        <span class="bold">Octotip:</span> You've activated the <em>file finder</em>
        by pressing <span class="kbd">t</span> Start typing to filter the
        file list. Use <span class="kbd badmono">↑</span> and
        <span class="kbd badmono">↓</span> to navigate,
        <span class="kbd">enter</span> to view files.
      </p>
    </div>

  <table class="tree-browser" cellpadding="0" cellspacing="0">
    <tr class="js-header"><th>&nbsp;</th><th>name</th></tr>
    <tr class="js-no-results no-results" style="display: none">
      <th colspan="2">No matching files</th>
    </tr>
    <tbody class="js-results-list js-navigation-container">
    </tbody>
  </table>
</div>

<div id="jump-to-line" style="display:none">
  <h2>Jump to Line</h2>
  <form accept-charset="UTF-8">
    <input name="utf8" type="hidden" value="&#x2713;" />
    <input class="textfield" type="text">
    <div class="full-button">
      <button type="submit" class="classy">
        <span>Go</span>
      </button>
    </div>
  </form>
</div>


<div class="subnav-bar">

  <ul class="actions subnav">
    <li><a href="/xushiwei/gobook/tags" class="blank" highlight="repo_tags">Tags <span class="counter">0</span></a></li>
    <li><a href="/xushiwei/gobook/downloads" class="blank downloads-blank" highlight="repo_downloads">Downloads <span class="counter">0</span></a></li>
    
  </ul>

  <ul class="scope">
    <li class="switcher">

      <div class="context-menu-container js-menu-container">
        <a href="#"
           class="minibutton bigger switcher js-menu-target js-commitish-button btn-branch repo-tree"
           data-master-branch="master"
           data-ref="master">
          <span><span class="icon"></span><i>branch:</i> master</span>
        </a>

        <div class="context-pane commitish-context js-menu-content">
          <a href="javascript:;" class="close js-menu-close"></a>
          <div class="context-title">Switch Branches/Tags</div>
          <div class="context-body pane-selector commitish-selector js-filterable-commitishes">
            <div class="filterbar">
              <input type="text" id="context-commitish-filter-field" class="commitish-filter" placeholder="Filter branches/tags" />

              <ul class="tabs">
                <li><a href="#" data-filter="branches" class="selected">Branches</a></li>
                <li><a href="#" data-filter="tags">Tags</a></li>
              </ul>
            </div>

              <div class="commitish-item branch-commitish selector-item">
                <h4>
                    <a href="/xushiwei/gobook/blob/master/dive-into/interface/02/interface.c" data-name="master" rel="nofollow">master</a>
                </h4>
              </div>


            <div class="no-results" style="display:none">Nothing to show</div>
          </div>
        </div><!-- /.commitish-context-context -->
      </div>

    </li>
  </ul>

  <ul class="subnav with-scope">

    <li><a href="/xushiwei/gobook" class="selected" highlight="repo_source">Files</a></li>
    <li><a href="/xushiwei/gobook/commits/master" highlight="repo_commits">Commits</a></li>
    <li><a href="/xushiwei/gobook/branches" class="" highlight="repo_branches" rel="nofollow">Branches <span class="counter">1</span></a></li>
  </ul>

</div>

  
  
  


          

        </div><!-- /.repohead -->

        





<!-- block_view_fragment_key: views7/v8/blob:v21:f852b513ea6b524ab6811a34d2868a9c -->
  <div id="slider">

    <div class="breadcrumb" data-path="dive-into/interface/02/interface.c/">
      <b itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/xushiwei/gobook/tree/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee" class="js-rewrite-sha" itemprop="url"><span itemprop="title">gobook</span></a></b> / <span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/xushiwei/gobook/tree/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into" class="js-rewrite-sha" itemscope="url"><span itemprop="title">dive-into</span></a></span> / <span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/xushiwei/gobook/tree/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into/interface" class="js-rewrite-sha" itemscope="url"><span itemprop="title">interface</span></a></span> / <span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/xushiwei/gobook/tree/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into/interface/02" class="js-rewrite-sha" itemscope="url"><span itemprop="title">02</span></a></span> / <strong class="final-path">interface.c</strong> <span class="js-clippy mini-icon clippy " data-clipboard-text="dive-into/interface/02/interface.c" data-copied-hint="copied!" data-copy-hint="copy to clipboard"></span>
    </div>


      <div class="commit file-history-tease" data-path="dive-into/interface/02/interface.c/">
        <img class="main-avatar" height="24" src="https://secure.gravatar.com/avatar/e85ffba4f3fb73edb6b58ebe12de7e9f?s=140&amp;d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png" width="24" />
        <span class="author"><a href="/xushiwei">xushiwei</a></span>
        <time class="js-relative-date" datetime="2012-04-15T02:26:36-07:00" title="2012-04-15 02:26:36">April 15, 2012</time>
        <div class="commit-title">
            <a href="/xushiwei/gobook/commit/fffed182e06d366ceab01a2924a41c8adf1c4072" class="message">dive into interface: query interface</a>
        </div>

        <div class="participation">
          <p class="quickstat"><a href="#blob_contributors_box" rel="facebox"><strong>1</strong> contributor</a></p>
          
        </div>
        <div id="blob_contributors_box" style="display:none">
          <h2>Users on GitHub who have contributed to this file</h2>
          <ul class="facebox-user-list">
            <li>
              <img height="24" src="https://secure.gravatar.com/avatar/e85ffba4f3fb73edb6b58ebe12de7e9f?s=140&amp;d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-140.png" width="24" />
              <a href="/xushiwei">xushiwei</a>
            </li>
          </ul>
        </div>
      </div>

    <div class="frames">
      <div class="frame frame-center" data-path="dive-into/interface/02/interface.c/" data-permalink-url="/xushiwei/gobook/blob/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into/interface/02/interface.c" data-title="gobook/dive-into/interface/02/interface.c at master · xushiwei/gobook · GitHub" data-type="blob">

        <div id="files" class="bubble">
          <div class="file">
            <div class="meta">
              <div class="info">
                <span class="icon"><b class="mini-icon text-file"></b></span>
                <span class="mode" title="File Mode">100644</span>
                  <span>140 lines (108 sloc)</span>
                <span>2.805 kb</span>
              </div>
              <ul class="button-group actions">
                  <li>
                    <a class="grouped-button file-edit-link minibutton bigger lighter js-rewrite-sha" href="/xushiwei/gobook/edit/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee/dive-into/interface/02/interface.c" data-method="post" rel="nofollow"><span>Edit this file</span></a>
                  </li>

                <li>
                  <a href="/xushiwei/gobook/raw/master/dive-into/interface/02/interface.c" class="minibutton btn-raw grouped-button bigger lighter" id="raw-url"><span><span class="icon"></span>Raw</span></a>
                </li>
                  <li>
                    <a href="/xushiwei/gobook/blame/master/dive-into/interface/02/interface.c" class="minibutton btn-blame grouped-button bigger lighter"><span><span class="icon"></span>Blame</span></a>
                  </li>
                <li>
                  <a href="/xushiwei/gobook/commits/master/dive-into/interface/02/interface.c" class="minibutton btn-history grouped-button bigger lighter" rel="nofollow"><span><span class="icon"></span>History</span></a>
                </li>
              </ul>
            </div>
              <div class="data type-c">
      <table cellpadding="0" cellspacing="0" class="lines">
        <tr>
          <td>
            <pre class="line_numbers"><span id="L1" rel="#L1">1</span>
<span id="L2" rel="#L2">2</span>
<span id="L3" rel="#L3">3</span>
<span id="L4" rel="#L4">4</span>
<span id="L5" rel="#L5">5</span>
<span id="L6" rel="#L6">6</span>
<span id="L7" rel="#L7">7</span>
<span id="L8" rel="#L8">8</span>
<span id="L9" rel="#L9">9</span>
<span id="L10" rel="#L10">10</span>
<span id="L11" rel="#L11">11</span>
<span id="L12" rel="#L12">12</span>
<span id="L13" rel="#L13">13</span>
<span id="L14" rel="#L14">14</span>
<span id="L15" rel="#L15">15</span>
<span id="L16" rel="#L16">16</span>
<span id="L17" rel="#L17">17</span>
<span id="L18" rel="#L18">18</span>
<span id="L19" rel="#L19">19</span>
<span id="L20" rel="#L20">20</span>
<span id="L21" rel="#L21">21</span>
<span id="L22" rel="#L22">22</span>
<span id="L23" rel="#L23">23</span>
<span id="L24" rel="#L24">24</span>
<span id="L25" rel="#L25">25</span>
<span id="L26" rel="#L26">26</span>
<span id="L27" rel="#L27">27</span>
<span id="L28" rel="#L28">28</span>
<span id="L29" rel="#L29">29</span>
<span id="L30" rel="#L30">30</span>
<span id="L31" rel="#L31">31</span>
<span id="L32" rel="#L32">32</span>
<span id="L33" rel="#L33">33</span>
<span id="L34" rel="#L34">34</span>
<span id="L35" rel="#L35">35</span>
<span id="L36" rel="#L36">36</span>
<span id="L37" rel="#L37">37</span>
<span id="L38" rel="#L38">38</span>
<span id="L39" rel="#L39">39</span>
<span id="L40" rel="#L40">40</span>
<span id="L41" rel="#L41">41</span>
<span id="L42" rel="#L42">42</span>
<span id="L43" rel="#L43">43</span>
<span id="L44" rel="#L44">44</span>
<span id="L45" rel="#L45">45</span>
<span id="L46" rel="#L46">46</span>
<span id="L47" rel="#L47">47</span>
<span id="L48" rel="#L48">48</span>
<span id="L49" rel="#L49">49</span>
<span id="L50" rel="#L50">50</span>
<span id="L51" rel="#L51">51</span>
<span id="L52" rel="#L52">52</span>
<span id="L53" rel="#L53">53</span>
<span id="L54" rel="#L54">54</span>
<span id="L55" rel="#L55">55</span>
<span id="L56" rel="#L56">56</span>
<span id="L57" rel="#L57">57</span>
<span id="L58" rel="#L58">58</span>
<span id="L59" rel="#L59">59</span>
<span id="L60" rel="#L60">60</span>
<span id="L61" rel="#L61">61</span>
<span id="L62" rel="#L62">62</span>
<span id="L63" rel="#L63">63</span>
<span id="L64" rel="#L64">64</span>
<span id="L65" rel="#L65">65</span>
<span id="L66" rel="#L66">66</span>
<span id="L67" rel="#L67">67</span>
<span id="L68" rel="#L68">68</span>
<span id="L69" rel="#L69">69</span>
<span id="L70" rel="#L70">70</span>
<span id="L71" rel="#L71">71</span>
<span id="L72" rel="#L72">72</span>
<span id="L73" rel="#L73">73</span>
<span id="L74" rel="#L74">74</span>
<span id="L75" rel="#L75">75</span>
<span id="L76" rel="#L76">76</span>
<span id="L77" rel="#L77">77</span>
<span id="L78" rel="#L78">78</span>
<span id="L79" rel="#L79">79</span>
<span id="L80" rel="#L80">80</span>
<span id="L81" rel="#L81">81</span>
<span id="L82" rel="#L82">82</span>
<span id="L83" rel="#L83">83</span>
<span id="L84" rel="#L84">84</span>
<span id="L85" rel="#L85">85</span>
<span id="L86" rel="#L86">86</span>
<span id="L87" rel="#L87">87</span>
<span id="L88" rel="#L88">88</span>
<span id="L89" rel="#L89">89</span>
<span id="L90" rel="#L90">90</span>
<span id="L91" rel="#L91">91</span>
<span id="L92" rel="#L92">92</span>
<span id="L93" rel="#L93">93</span>
<span id="L94" rel="#L94">94</span>
<span id="L95" rel="#L95">95</span>
<span id="L96" rel="#L96">96</span>
<span id="L97" rel="#L97">97</span>
<span id="L98" rel="#L98">98</span>
<span id="L99" rel="#L99">99</span>
<span id="L100" rel="#L100">100</span>
<span id="L101" rel="#L101">101</span>
<span id="L102" rel="#L102">102</span>
<span id="L103" rel="#L103">103</span>
<span id="L104" rel="#L104">104</span>
<span id="L105" rel="#L105">105</span>
<span id="L106" rel="#L106">106</span>
<span id="L107" rel="#L107">107</span>
<span id="L108" rel="#L108">108</span>
<span id="L109" rel="#L109">109</span>
<span id="L110" rel="#L110">110</span>
<span id="L111" rel="#L111">111</span>
<span id="L112" rel="#L112">112</span>
<span id="L113" rel="#L113">113</span>
<span id="L114" rel="#L114">114</span>
<span id="L115" rel="#L115">115</span>
<span id="L116" rel="#L116">116</span>
<span id="L117" rel="#L117">117</span>
<span id="L118" rel="#L118">118</span>
<span id="L119" rel="#L119">119</span>
<span id="L120" rel="#L120">120</span>
<span id="L121" rel="#L121">121</span>
<span id="L122" rel="#L122">122</span>
<span id="L123" rel="#L123">123</span>
<span id="L124" rel="#L124">124</span>
<span id="L125" rel="#L125">125</span>
<span id="L126" rel="#L126">126</span>
<span id="L127" rel="#L127">127</span>
<span id="L128" rel="#L128">128</span>
<span id="L129" rel="#L129">129</span>
<span id="L130" rel="#L130">130</span>
<span id="L131" rel="#L131">131</span>
<span id="L132" rel="#L132">132</span>
<span id="L133" rel="#L133">133</span>
<span id="L134" rel="#L134">134</span>
<span id="L135" rel="#L135">135</span>
<span id="L136" rel="#L136">136</span>
<span id="L137" rel="#L137">137</span>
<span id="L138" rel="#L138">138</span>
<span id="L139" rel="#L139">139</span>
</pre>
          </td>
          <td width="100%">
                <div class="highlight"><pre><div class='line' id='LC1'><span class="cp">#include &lt;stdio.h&gt;</span></div><div class='line' id='LC2'><span class="cp">#include &lt;stdlib.h&gt;</span></div><div class='line' id='LC3'><br/></div><div class='line' id='LC4'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC5'><br/></div><div class='line' id='LC6'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_TypeInfo</span> <span class="p">{</span></div><div class='line' id='LC7'>	<span class="c1">// 用于运行时取得类型信息, 比如反射机制</span></div><div class='line' id='LC8'><span class="p">}</span> <span class="n">TypeInfo</span><span class="p">;</span></div><div class='line' id='LC9'><br/></div><div class='line' id='LC10'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_InterfaceInfo</span> <span class="p">{</span></div><div class='line' id='LC11'>	<span class="c1">// 用于运行时取得interface信息</span></div><div class='line' id='LC12'><span class="p">}</span> <span class="n">InterfaceInfo</span><span class="p">;</span></div><div class='line' id='LC13'><br/></div><div class='line' id='LC14'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC15'><br/></div><div class='line' id='LC16'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_IReadWriterTbl</span> <span class="p">{</span></div><div class='line' id='LC17'>	<span class="n">InterfaceInfo</span><span class="o">*</span> <span class="n">inter</span><span class="p">;</span></div><div class='line' id='LC18'>	<span class="n">TypeInfo</span><span class="o">*</span> <span class="n">type</span><span class="p">;</span></div><div class='line' id='LC19'>	<span class="kt">int</span> <span class="p">(</span><span class="o">*</span><span class="n">Read</span><span class="p">)(</span><span class="kt">void</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">);</span></div><div class='line' id='LC20'>	<span class="kt">int</span> <span class="p">(</span><span class="o">*</span><span class="n">Write</span><span class="p">)(</span><span class="kt">void</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">);</span></div><div class='line' id='LC21'><span class="p">}</span> <span class="n">IReadWriterTbl</span><span class="p">;</span></div><div class='line' id='LC22'><br/></div><div class='line' id='LC23'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_IReadWriter</span> <span class="p">{</span></div><div class='line' id='LC24'>	<span class="n">IReadWriterTbl</span><span class="o">*</span> <span class="n">tab</span><span class="p">;</span></div><div class='line' id='LC25'>	<span class="kt">void</span><span class="o">*</span> <span class="n">data</span><span class="p">;</span></div><div class='line' id='LC26'><span class="p">}</span> <span class="n">IReadWriter</span><span class="p">;</span></div><div class='line' id='LC27'><br/></div><div class='line' id='LC28'><span class="n">InterfaceInfo</span> <span class="n">g_InterfaceInfo_IReadWriter</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC29'>	<span class="c1">// ...</span></div><div class='line' id='LC30'><span class="p">};</span></div><div class='line' id='LC31'><br/></div><div class='line' id='LC32'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC33'><br/></div><div class='line' id='LC34'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_IWriterTbl</span> <span class="p">{</span></div><div class='line' id='LC35'>	<span class="n">InterfaceInfo</span><span class="o">*</span> <span class="n">inter</span><span class="p">;</span></div><div class='line' id='LC36'>	<span class="n">TypeInfo</span><span class="o">*</span> <span class="n">type</span><span class="p">;</span></div><div class='line' id='LC37'>	<span class="kt">int</span> <span class="p">(</span><span class="o">*</span><span class="n">Write</span><span class="p">)(</span><span class="kt">void</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">);</span></div><div class='line' id='LC38'><span class="p">}</span> <span class="n">IWriterTbl</span><span class="p">;</span></div><div class='line' id='LC39'><br/></div><div class='line' id='LC40'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_IWriter</span> <span class="p">{</span></div><div class='line' id='LC41'>	<span class="n">IWriterTbl</span><span class="o">*</span> <span class="n">tab</span><span class="p">;</span></div><div class='line' id='LC42'>	<span class="kt">void</span><span class="o">*</span> <span class="n">data</span><span class="p">;</span></div><div class='line' id='LC43'><span class="p">}</span> <span class="n">IWriter</span><span class="p">;</span></div><div class='line' id='LC44'><br/></div><div class='line' id='LC45'><span class="n">InterfaceInfo</span> <span class="n">g_InterfaceInfo_IWriter</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC46'>	<span class="c1">// ...</span></div><div class='line' id='LC47'><span class="p">};</span></div><div class='line' id='LC48'><br/></div><div class='line' id='LC49'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC50'><br/></div><div class='line' id='LC51'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_A</span> <span class="p">{</span></div><div class='line' id='LC52'>	<span class="kt">int</span> <span class="n">a</span><span class="p">;</span></div><div class='line' id='LC53'><span class="p">}</span> <span class="n">A</span><span class="p">;</span></div><div class='line' id='LC54'><br/></div><div class='line' id='LC55'><span class="kt">int</span> <span class="nf">A_Read</span><span class="p">(</span><span class="n">A</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC56'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;A_Read: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">this</span><span class="o">-&gt;</span><span class="n">a</span><span class="p">);</span></div><div class='line' id='LC57'>	<span class="k">return</span> <span class="n">cb</span><span class="p">;</span></div><div class='line' id='LC58'><span class="p">}</span></div><div class='line' id='LC59'><br/></div><div class='line' id='LC60'><span class="kt">int</span> <span class="nf">A_Write</span><span class="p">(</span><span class="n">A</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC61'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;A_Write: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">this</span><span class="o">-&gt;</span><span class="n">a</span><span class="p">);</span></div><div class='line' id='LC62'>	<span class="k">return</span> <span class="n">cb</span><span class="p">;</span></div><div class='line' id='LC63'><span class="p">}</span></div><div class='line' id='LC64'><br/></div><div class='line' id='LC65'><span class="n">TypeInfo</span> <span class="n">g_TypeInfo_A</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC66'>	<span class="c1">// ...</span></div><div class='line' id='LC67'><span class="p">};</span></div><div class='line' id='LC68'><br/></div><div class='line' id='LC69'><span class="n">A</span><span class="o">*</span> <span class="nf">NewA</span><span class="p">(</span><span class="kt">int</span> <span class="n">params</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC70'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;NewA: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">params</span><span class="p">);</span></div><div class='line' id='LC71'>	<span class="n">A</span><span class="o">*</span> <span class="n">this</span> <span class="o">=</span> <span class="p">(</span><span class="n">A</span><span class="o">*</span><span class="p">)</span><span class="n">malloc</span><span class="p">(</span><span class="k">sizeof</span><span class="p">(</span><span class="n">A</span><span class="p">));</span></div><div class='line' id='LC72'>	<span class="n">this</span><span class="o">-&gt;</span><span class="n">a</span> <span class="o">=</span> <span class="n">params</span><span class="p">;</span></div><div class='line' id='LC73'>	<span class="k">return</span> <span class="n">this</span><span class="p">;</span></div><div class='line' id='LC74'><span class="p">}</span></div><div class='line' id='LC75'><br/></div><div class='line' id='LC76'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC77'><br/></div><div class='line' id='LC78'><span class="k">typedef</span> <span class="k">struct</span> <span class="n">_B</span> <span class="p">{</span></div><div class='line' id='LC79'>	<span class="n">A</span> <span class="n">base</span><span class="p">;</span></div><div class='line' id='LC80'><span class="p">}</span> <span class="n">B</span><span class="p">;</span></div><div class='line' id='LC81'><br/></div><div class='line' id='LC82'><span class="kt">int</span> <span class="nf">B_Write</span><span class="p">(</span><span class="n">B</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC83'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;B_Write: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">this</span><span class="o">-&gt;</span><span class="n">base</span><span class="p">.</span><span class="n">a</span><span class="p">);</span></div><div class='line' id='LC84'>	<span class="k">return</span> <span class="n">cb</span><span class="p">;</span></div><div class='line' id='LC85'><span class="p">}</span></div><div class='line' id='LC86'><br/></div><div class='line' id='LC87'><span class="kt">void</span> <span class="nf">B_Foo</span><span class="p">(</span><span class="n">B</span><span class="o">*</span> <span class="n">this</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC88'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;B_Foo: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">this</span><span class="o">-&gt;</span><span class="n">base</span><span class="p">.</span><span class="n">a</span><span class="p">);</span></div><div class='line' id='LC89'><span class="p">}</span></div><div class='line' id='LC90'><br/></div><div class='line' id='LC91'><span class="n">TypeInfo</span> <span class="n">g_TypeInfo_B</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC92'>	<span class="c1">// ...</span></div><div class='line' id='LC93'><span class="p">};</span></div><div class='line' id='LC94'><br/></div><div class='line' id='LC95'><span class="n">B</span><span class="o">*</span> <span class="nf">NewB</span><span class="p">(</span><span class="kt">int</span> <span class="n">params</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC96'>	<span class="n">printf</span><span class="p">(</span><span class="s">&quot;NewB: %d</span><span class="se">\n</span><span class="s">&quot;</span><span class="p">,</span> <span class="n">params</span><span class="p">);</span></div><div class='line' id='LC97'>	<span class="n">B</span><span class="o">*</span> <span class="n">this</span> <span class="o">=</span> <span class="p">(</span><span class="n">B</span><span class="o">*</span><span class="p">)</span><span class="n">malloc</span><span class="p">(</span><span class="k">sizeof</span><span class="p">(</span><span class="n">B</span><span class="p">));</span></div><div class='line' id='LC98'>	<span class="n">this</span><span class="o">-&gt;</span><span class="n">base</span><span class="p">.</span><span class="n">a</span> <span class="o">=</span> <span class="n">params</span><span class="p">;</span></div><div class='line' id='LC99'>	<span class="k">return</span> <span class="n">this</span><span class="p">;</span></div><div class='line' id='LC100'><span class="p">}</span></div><div class='line' id='LC101'><br/></div><div class='line' id='LC102'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC103'><br/></div><div class='line' id='LC104'><span class="n">IReadWriterTbl</span> <span class="n">g_Itbl_IReadWriter_B</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC105'>	<span class="o">&amp;</span><span class="n">g_InterfaceInfo_IReadWriter</span><span class="p">,</span></div><div class='line' id='LC106'>	<span class="o">&amp;</span><span class="n">g_TypeInfo_B</span><span class="p">,</span></div><div class='line' id='LC107'>	<span class="p">(</span><span class="kt">int</span> <span class="p">(</span><span class="o">*</span><span class="p">)(</span><span class="kt">void</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">))</span><span class="n">A_Read</span><span class="p">,</span></div><div class='line' id='LC108'>	<span class="p">(</span><span class="kt">int</span> <span class="p">(</span><span class="o">*</span><span class="p">)(</span><span class="kt">void</span><span class="o">*</span> <span class="n">this</span><span class="p">,</span> <span class="kt">char</span><span class="o">*</span> <span class="n">buf</span><span class="p">,</span> <span class="kt">int</span> <span class="n">cb</span><span class="p">))</span><span class="n">B_Write</span></div><div class='line' id='LC109'><span class="p">};</span></div><div class='line' id='LC110'><br/></div><div class='line' id='LC111'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC112'><br/></div><div class='line' id='LC113'><span class="n">IWriterTbl</span><span class="o">*</span> <span class="nf">Itbl_IWriter_From_IReadWriter</span><span class="p">(</span><span class="n">IReadWriterTbl</span><span class="o">*</span> <span class="n">src</span><span class="p">)</span> <span class="p">{</span></div><div class='line' id='LC114'>	<span class="n">IWriterTbl</span><span class="o">*</span> <span class="n">dest</span> <span class="o">=</span> <span class="p">(</span><span class="n">IWriterTbl</span><span class="o">*</span><span class="p">)</span><span class="n">malloc</span><span class="p">(</span><span class="k">sizeof</span><span class="p">(</span><span class="n">IWriterTbl</span><span class="p">));</span></div><div class='line' id='LC115'>	<span class="n">dest</span><span class="o">-&gt;</span><span class="n">inter</span> <span class="o">=</span> <span class="o">&amp;</span><span class="n">g_InterfaceInfo_IWriter</span><span class="p">,</span></div><div class='line' id='LC116'>	<span class="n">dest</span><span class="o">-&gt;</span><span class="n">type</span> <span class="o">=</span> <span class="n">src</span><span class="o">-&gt;</span><span class="n">type</span><span class="p">;</span></div><div class='line' id='LC117'>	<span class="n">dest</span><span class="o">-&gt;</span><span class="n">Write</span> <span class="o">=</span> <span class="n">src</span><span class="o">-&gt;</span><span class="n">Write</span><span class="p">;</span></div><div class='line' id='LC118'>	<span class="k">return</span> <span class="n">dest</span><span class="p">;</span></div><div class='line' id='LC119'><span class="p">}</span></div><div class='line' id='LC120'><br/></div><div class='line' id='LC121'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC122'><br/></div><div class='line' id='LC123'><span class="kt">int</span> <span class="nf">main</span><span class="p">()</span> <span class="p">{</span></div><div class='line' id='LC124'>	<span class="n">B</span><span class="o">*</span> <span class="n">unnamed</span> <span class="o">=</span> <span class="n">NewB</span><span class="p">(</span><span class="mi">9</span><span class="p">);</span></div><div class='line' id='LC125'>	<span class="n">IReadWriter</span> <span class="n">p</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC126'>		<span class="o">&amp;</span><span class="n">g_Itbl_IReadWriter_B</span><span class="p">,</span></div><div class='line' id='LC127'>		<span class="n">unnamed</span></div><div class='line' id='LC128'>	<span class="p">};</span></div><div class='line' id='LC129'>	<span class="n">IWriter</span> <span class="n">p2</span> <span class="o">=</span> <span class="p">{</span></div><div class='line' id='LC130'>		<span class="n">Itbl_IWriter_From_IReadWriter</span><span class="p">(</span><span class="n">p</span><span class="p">.</span><span class="n">tab</span><span class="p">),</span></div><div class='line' id='LC131'>		<span class="n">p</span><span class="p">.</span><span class="n">data</span></div><div class='line' id='LC132'>	<span class="p">};</span></div><div class='line' id='LC133'>	<span class="n">p</span><span class="p">.</span><span class="n">tab</span><span class="o">-&gt;</span><span class="n">Read</span><span class="p">(</span><span class="n">p</span><span class="p">.</span><span class="n">data</span><span class="p">,</span> <span class="nb">NULL</span><span class="p">,</span> <span class="mi">10</span><span class="p">);</span></div><div class='line' id='LC134'>	<span class="n">p2</span><span class="p">.</span><span class="n">tab</span><span class="o">-&gt;</span><span class="n">Write</span><span class="p">(</span><span class="n">p2</span><span class="p">.</span><span class="n">data</span><span class="p">,</span> <span class="nb">NULL</span><span class="p">,</span> <span class="mi">10</span><span class="p">);</span></div><div class='line' id='LC135'>	<span class="k">return</span> <span class="mi">0</span><span class="p">;</span></div><div class='line' id='LC136'><span class="p">}</span></div><div class='line' id='LC137'><br/></div><div class='line' id='LC138'><span class="c1">// -------------------------------------------------------------</span></div><div class='line' id='LC139'><br/></div></pre></div>
          </td>
        </tr>
      </table>
  </div>

          </div>
        </div>
      </div>
    </div>

  </div>

<div class="frame frame-loading large-loading-area" style="display:none;" data-tree-list-url="/xushiwei/gobook/tree-list/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee" data-blob-url-prefix="/xushiwei/gobook/blob/e5cbc33e10ee41b35ac5fd4e3ec33dd20f4abaee">
  <img src="https://a248.e.akamai.net/assets.github.com/images/spinners/octocat-spinner-64.gif?1329872006" height="64" width="64">
</div>

      </div>
      <div class="context-overlay"></div>
    </div>

      <div id="footer-push"></div><!-- hack for sticky footer -->
    </div><!-- end of wrapper - hack for sticky footer -->

      <!-- footer -->
      <div id="footer" >
        
  <div class="upper_footer">
     <div class="container clearfix">

       <!--[if IE]><h4 id="blacktocat_ie">GitHub Links</h4><![endif]-->
       <![if !IE]><h4 id="blacktocat">GitHub Links</h4><![endif]>

       <ul class="footer_nav">
         <h4>GitHub</h4>
         <li><a href="https://github.com/about">About</a></li>
         <li><a href="https://github.com/blog">Blog</a></li>
         <li><a href="https://github.com/features">Features</a></li>
         <li><a href="https://github.com/contact">Contact &amp; Support</a></li>
         <li><a href="https://github.com/training">Training</a></li>
         <li><a href="http://enterprise.github.com/">GitHub Enterprise</a></li>
         <li><a href="http://status.github.com/">Site Status</a></li>
       </ul>

       <ul class="footer_nav">
         <h4>Tools</h4>
         <li><a href="http://get.gaug.es/">Gauges: Analyze web traffic</a></li>
         <li><a href="http://speakerdeck.com">Speaker Deck: Presentations</a></li>
         <li><a href="https://gist.github.com">Gist: Code snippets</a></li>
         <li><a href="http://mac.github.com/">GitHub for Mac</a></li>
         <li><a href="http://mobile.github.com/">Issues for iPhone</a></li>
         <li><a href="http://jobs.github.com/">Job Board</a></li>
       </ul>

       <ul class="footer_nav">
         <h4>Extras</h4>
         <li><a href="http://shop.github.com/">GitHub Shop</a></li>
         <li><a href="http://octodex.github.com/">The Octodex</a></li>
       </ul>

       <ul class="footer_nav">
         <h4>Documentation</h4>
         <li><a href="http://help.github.com/">GitHub Help</a></li>
         <li><a href="http://developer.github.com/">Developer API</a></li>
         <li><a href="http://github.github.com/github-flavored-markdown/">GitHub Flavored Markdown</a></li>
         <li><a href="http://pages.github.com/">GitHub Pages</a></li>
       </ul>

     </div><!-- /.site -->
  </div><!-- /.upper_footer -->

<div class="lower_footer">
  <div class="container clearfix">
    <!--[if IE]><div id="legal_ie"><![endif]-->
    <![if !IE]><div id="legal"><![endif]>
      <ul>
          <li><a href="https://github.com/site/terms">Terms of Service</a></li>
          <li><a href="https://github.com/site/privacy">Privacy</a></li>
          <li><a href="https://github.com/security">Security</a></li>
      </ul>

      <p>&copy; 2012 <span title="0.17301s from fe3.rs.github.com">GitHub</span> Inc. All rights reserved.</p>
    </div><!-- /#legal or /#legal_ie-->

      <div class="sponsor">
        <a href="http://www.rackspace.com" class="logo">
          <img alt="Dedicated Server" height="36" src="https://a248.e.akamai.net/assets.github.com/images/modules/footer/rackspaces_logo.png?1329521041" width="38" />
        </a>
        Powered by the <a href="http://www.rackspace.com ">Dedicated
        Servers</a> and<br/> <a href="http://www.rackspacecloud.com">Cloud
        Computing</a> of Rackspace Hosting<span>&reg;</span>
      </div>
  </div><!-- /.site -->
</div><!-- /.lower_footer -->

      </div><!-- /#footer -->

    

<div id="keyboard_shortcuts_pane" class="instapaper_ignore readability-extra" style="display:none">
  <h2>Keyboard Shortcuts <small><a href="#" class="js-see-all-keyboard-shortcuts">(see all)</a></small></h2>

  <div class="columns threecols">
    <div class="column first">
      <h3>Site wide shortcuts</h3>
      <dl class="keyboard-mappings">
        <dt>s</dt>
        <dd>Focus site search</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>?</dt>
        <dd>Bring up this help dialog</dd>
      </dl>
    </div><!-- /.column.first -->

    <div class="column middle" style='display:none'>
      <h3>Commit list</h3>
      <dl class="keyboard-mappings">
        <dt>j</dt>
        <dd>Move selection down</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>k</dt>
        <dd>Move selection up</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>c <em>or</em> o <em>or</em> enter</dt>
        <dd>Open commit</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>y</dt>
        <dd>Expand URL to its canonical form</dd>
      </dl>
    </div><!-- /.column.first -->

    <div class="column last" style='display:none'>
      <h3>Pull request list</h3>
      <dl class="keyboard-mappings">
        <dt>j</dt>
        <dd>Move selection down</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>k</dt>
        <dd>Move selection up</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt>o <em>or</em> enter</dt>
        <dd>Open issue</dd>
      </dl>
      <dl class="keyboard-mappings">
        <dt><span class="platform-mac">⌘</span><span class="platform-other">ctrl</span> <em>+</em> enter</dt>
        <dd>Submit comment</dd>
      </dl>
    </div><!-- /.columns.last -->

  </div><!-- /.columns.equacols -->

  <div style='display:none'>
    <div class="rule"></div>

    <h3>Issues</h3>

    <div class="columns threecols">
      <div class="column first">
        <dl class="keyboard-mappings">
          <dt>j</dt>
          <dd>Move selection down</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>k</dt>
          <dd>Move selection up</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>x</dt>
          <dd>Toggle selection</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>o <em>or</em> enter</dt>
          <dd>Open issue</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt><span class="platform-mac">⌘</span><span class="platform-other">ctrl</span> <em>+</em> enter</dt>
          <dd>Submit comment</dd>
        </dl>
      </div><!-- /.column.first -->
      <div class="column last">
        <dl class="keyboard-mappings">
          <dt>c</dt>
          <dd>Create issue</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>l</dt>
          <dd>Create label</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>i</dt>
          <dd>Back to inbox</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>u</dt>
          <dd>Back to issues</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>/</dt>
          <dd>Focus issues search</dd>
        </dl>
      </div>
    </div>
  </div>

  <div style='display:none'>
    <div class="rule"></div>

    <h3>Issues Dashboard</h3>

    <div class="columns threecols">
      <div class="column first">
        <dl class="keyboard-mappings">
          <dt>j</dt>
          <dd>Move selection down</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>k</dt>
          <dd>Move selection up</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>o <em>or</em> enter</dt>
          <dd>Open issue</dd>
        </dl>
      </div><!-- /.column.first -->
    </div>
  </div>

  <div style='display:none'>
    <div class="rule"></div>

    <h3>Network Graph</h3>
    <div class="columns equacols">
      <div class="column first">
        <dl class="keyboard-mappings">
          <dt><span class="badmono">←</span> <em>or</em> h</dt>
          <dd>Scroll left</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt><span class="badmono">→</span> <em>or</em> l</dt>
          <dd>Scroll right</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt><span class="badmono">↑</span> <em>or</em> k</dt>
          <dd>Scroll up</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt><span class="badmono">↓</span> <em>or</em> j</dt>
          <dd>Scroll down</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>t</dt>
          <dd>Toggle visibility of head labels</dd>
        </dl>
      </div><!-- /.column.first -->
      <div class="column last">
        <dl class="keyboard-mappings">
          <dt>shift <span class="badmono">←</span> <em>or</em> shift h</dt>
          <dd>Scroll all the way left</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>shift <span class="badmono">→</span> <em>or</em> shift l</dt>
          <dd>Scroll all the way right</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>shift <span class="badmono">↑</span> <em>or</em> shift k</dt>
          <dd>Scroll all the way up</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>shift <span class="badmono">↓</span> <em>or</em> shift j</dt>
          <dd>Scroll all the way down</dd>
        </dl>
      </div><!-- /.column.last -->
    </div>
  </div>

  <div >
    <div class="rule"></div>
    <div class="columns threecols">
      <div class="column first" >
        <h3>Source Code Browsing</h3>
        <dl class="keyboard-mappings">
          <dt>t</dt>
          <dd>Activates the file finder</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>l</dt>
          <dd>Jump to line</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>w</dt>
          <dd>Switch branch/tag</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>y</dt>
          <dd>Expand URL to its canonical form</dd>
        </dl>
      </div>
    </div>
  </div>

  <div style='display:none'>
    <div class="rule"></div>
    <div class="columns threecols">
      <div class="column first">
        <h3>Browsing Commits</h3>
        <dl class="keyboard-mappings">
          <dt><span class="platform-mac">⌘</span><span class="platform-other">ctrl</span> <em>+</em> enter</dt>
          <dd>Submit comment</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>escape</dt>
          <dd>Close form</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>p</dt>
          <dd>Parent commit</dd>
        </dl>
        <dl class="keyboard-mappings">
          <dt>o</dt>
          <dd>Other parent commit</dd>
        </dl>
      </div>
    </div>
  </div>
</div>

    <div id="markdown-help" class="instapaper_ignore readability-extra">
  <h2>Markdown Cheat Sheet</h2>

  <div class="cheatsheet-content">

  <div class="mod">
    <div class="col">
      <h3>Format Text</h3>
      <p>Headers</p>
      <pre>
# This is an &lt;h1&gt; tag
## This is an &lt;h2&gt; tag
###### This is an &lt;h6&gt; tag</pre>
     <p>Text styles</p>
     <pre>
*This text will be italic*
_This will also be italic_
**This text will be bold**
__This will also be bold__

*You **can** combine them*
</pre>
    </div>
    <div class="col">
      <h3>Lists</h3>
      <p>Unordered</p>
      <pre>
* Item 1
* Item 2
  * Item 2a
  * Item 2b</pre>
     <p>Ordered</p>
     <pre>
1. Item 1
2. Item 2
3. Item 3
   * Item 3a
   * Item 3b</pre>
    </div>
    <div class="col">
      <h3>Miscellaneous</h3>
      <p>Images</p>
      <pre>
![GitHub Logo](/images/logo.png)
Format: ![Alt Text](url)
</pre>
     <p>Links</p>
     <pre>
http://github.com - automatic!
[GitHub](http://github.com)</pre>
<p>Blockquotes</p>
     <pre>
As Kanye West said:

> We're living the future so
> the present is our past.
</pre>
    </div>
  </div>
  <div class="rule"></div>

  <h3>Code Examples in Markdown</h3>
  <div class="col">
      <p>Syntax highlighting with <a href="http://github.github.com/github-flavored-markdown/" title="GitHub Flavored Markdown" target="_blank">GFM</a></p>
      <pre>
```javascript
function fancyAlert(arg) {
  if(arg) {
    $.facebox({div:'#foo'})
  }
}
```</pre>
    </div>
    <div class="col">
      <p>Or, indent your code 4 spaces</p>
      <pre>
Here is a Python code example
without syntax highlighting:

    def foo:
      if not bar:
        return true</pre>
    </div>
    <div class="col">
      <p>Inline code for comments</p>
      <pre>
I think you should use an
`&lt;addr&gt;` element here instead.</pre>
    </div>
  </div>

  </div>
</div>


    <div class="ajax-error-message">
      <p><span class="icon"></span> Something went wrong with that request. Please try again. <a href="javascript:;" class="ajax-error-dismiss">Dismiss</a></p>
    </div>

    <div id="logo-popup">
      <h2>Looking for the GitHub logo?</h2>
      <ul>
        <li>
          <h4>GitHub Logo</h4>
          <a href="http://github-media-downloads.s3.amazonaws.com/GitHub_Logos.zip"><img alt="Github_logo" src="https://a248.e.akamai.net/assets.github.com/images/modules/about_page/github_logo.png?1306884374" /></a>
          <a href="http://github-media-downloads.s3.amazonaws.com/GitHub_Logos.zip" class="minibutton btn-download download"><span><span class="icon"></span>Download</span></a>
        </li>
        <li>
          <h4>The Octocat</h4>
          <a href="http://github-media-downloads.s3.amazonaws.com/Octocats.zip"><img alt="Octocat" src="https://a248.e.akamai.net/assets.github.com/images/modules/about_page/octocat.png?1306884374" /></a>
          <a href="http://github-media-downloads.s3.amazonaws.com/Octocats.zip" class="minibutton btn-download download"><span><span class="icon"></span>Download</span></a>
        </li>
      </ul>
    </div>

    
    
    
    <span id='server_response_time' data-time='0.17556' data-host='fe3'></span>
  </body>
</html>

