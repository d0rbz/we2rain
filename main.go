package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{}

func main() {

	channels := []string{
		"https://t.me/s/we2rain",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/iran_v2ray1",
		"https://t.me/s/v2rayngn",
		"https://t.me/s/Alfred_Config",
		"https://t.me/s/freakConfig",
		"https://t.me/s/v2ray_Alpha",
		"https://t.me/s/v2ray_configs_pool",
		"https://t.me/vpnfail_v2ray",
		"https://t.me/ConfigsHUB",
		"https://t.me/VlessVpnFree",
		"https://t.me/vpnfail_vless",
		"https://t.me/ev2rayy",
		"https://t.me/DailyV2RY",
		"https://t.me/IP_CF_Config",
		"https://t.me/BigSmoke_Config",
		"https://t.me/configMs",
		"https://t.me/s/v2rayng_fa2",
		"https://t.me/s/v2rayng_org",
		"https://t.me/s/V2rayNGvpni",
		"https://t.me/s/custom_14",
		"https://t.me/s/v2rayNG_VPNN",
		"https://t.me/s/v2ray_outlineir",
		"https://t.me/s/v2_vmess",
		"https://t.me/s/FreeVlessVpn",
		"https://t.me/s/freeland8",
		"https://t.me/s/vmess_vless_v2rayng",
		"https://t.me/s/PrivateVPNs",
		"https://t.me/s/vmessiran",
		"https://t.me/s/Outline_Vpn",
		"https://t.me/s/vmessq",
		"https://t.me/s/WeePeeN",
		"https://t.me/s/V2rayNG3",
		"https://t.me/s/ShadowsocksM",
		"https://t.me/s/shadowsocksshop",
		"https://t.me/s/v2rayan",
		"https://t.me/ShadowSocks_s",
		"https://t.me/s/VmessProtocol",
		"https://t.me/s/napsternetv_config",
		"https://t.me/s/Easy_Free_VPN",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/V2RAY_VMESS_free",
		"https://t.me/s/v2ray_for_free",
		"https://t.me/s/V2rayN_Free",
		"https://t.me/s/free4allVPN",
		"https://t.me/s/vpn_ocean",
		"https://t.me/s/configV2rayForFree",
		"https://t.me/s/FreeV2rays",
		"https://t.me/s/DigiV2ray",
		"https://t.me/s/v2rayNG_VPN",
		"https://t.me/s/freev2rayssr",
		"https://t.me/s/v2rayn_server",
		"https://t.me/s/Shadowlinkserverr",
		"https://t.me/s/iranvpnet",
		"https://t.me/s/vmess_iran",
		"https://t.me/s/mahsaamoon1",
		"https://t.me/s/V2RAY_NEW",
		"https://t.me/s/v2RayChannel",
		"https://t.me/s/configV2rayNG",
		"https://t.me/s/config_v2ray",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/vpnmasi",
		"https://t.me/s/v2ray_custom",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/HTTPCustomLand",
		"https://t.me/s/ViPVpn_v2ray",
		"https://t.me/s/FreeNet1500",
		"https://t.me/s/v2ray_ar",
		"https://t.me/s/beta_v2ray",
		"https://t.me/s/vip_vpn_2022",
		"https://t.me/s/FOX_VPN66",
		"https://t.me/s/VorTexIRN",
		"https://t.me/s/YtTe3la",
		"https://t.me/s/V2RayOxygen",
		"https://t.me/s/Network_442",
		"https://t.me/s/VPN_443",
		"https://t.me/s/v2rayng_v",
		"https://t.me/s/ultrasurf_12",
		"https://t.me/V2rayNGn",
		"https://t.me/free4allVPN",
		"https://t.me/PrivateVPNs",
		"https://t.me/V2rayng_Fast",
		"https://t.me/DirectVPN",
		"https://t.me/ProxyFn",
		"https://t.me/v2ray_outlineir",
		"https://t.me/NetAccount",
		"https://t.me/oneclickvpnkeys",
		"https://t.me/daorzadannet",
		"https://t.me/LoRd_uL4mo",
		"https://t.me/Outline_Vpn",
		"https://t.me/vpn_xw",
		"https://t.me/prrofile_purple",
		"https://t.me/proxyymeliii",
		"https://t.me/azadi_az_inja_migzare",
		"https://t.me/WomanLifeFreedomVPN",
		"https://t.me/MsV2ray",
		"https://t.me/internet4iran",
		"https://t.me/LegenderY_Servers",
		"https://t.me/vpnfail_v2ray",
		"https://t.me/free_v2rayy",
		"https://t.me/UnlimitedDev",
		"https://t.me/vmessorg",
		"https://t.me/v2rayNG_Matsuri",
		"https://t.me/v2ray1_ng",
		"https://t.me/v2rayngvpn",
		"https://t.me/vpn_ioss",
		"https://t.me/v2freevpn",
		"https://t.me/vless_vmess",
		"https://t.me/customv2ray",
		"https://t.me/FalconPolV2rayNG",
		"https://t.me/Jeyksatan",
		"https://t.me/MTConfig",
		"https://t.me/hassan_saboorii",
		"https://t.me/v2rayvmess",
		"https://t.me/v2rayNGNeT",
		"https://t.me/lagvpn13",
		"https://t.me/server01012",
		"https://t.me/ShadowProxy66",
		"https://t.me/ipV2Ray",
		"https://t.me/v2rayNG_VPNN",
		"https://t.me/kiava",
		"https://t.me/Helix_Servers",
		"https://t.me/PAINB0Y",
		"https://t.me/vmess_vless_v2rayng",
		"https://t.me/VpnProSec",
		"https://t.me/VlessConfig",
		"https://t.me/NIM_VPN_ir",
		"https://t.me/FreeIranT",
		"https://t.me/hashmakvpn",
		"https://t.me/X_Her0",
		"https://t.me/Napsternetvirani",
		"https://t.me/Cov2ray",
		"https://t.me/iran_ray",
		"https://t.me/INIT1984",
		"https://t.me/EXOGAMERS",
		"https://t.me/V2RayTz",
		"https://t.me/ServerNett",
		"https://t.me/Pinkpotatocloud",
		"https://t.me/CloudCityy",
		"https://t.me/VmessProtocol",
		"https://t.me/DarkVPNpro",
		"https://t.me/Qv2raychannel",
		"https://t.me/xrayzxn",
		"https://t.me/MehradLearn",
		"https://t.me/shopingv2ray",
		"https://t.me/xrayproxy",
		"https://t.me/Proxy_PJ",
		"https://t.me/SafeNet_Server",
		"https://t.me/FreeV2rays",
		"https://t.me/configV2rayNG",
		"https://t.me/vpnmasi",
		"https://t.me/v2ray_ar",
		"https://t.me/iSeqaro",
		"https://t.me/God_CONFIG",
		"https://t.me/vpn_ioss",
		"https://t.me/AliAlma_GSM",
		"https://t.me/reality_daily",
		"https://t.me/v2rayng_org",
		"https://t.me/v2rayngvpn",
		"https://t.me/flyv2ray",
		"https://t.me/v2ray_outlineir",
		"https://t.me/v2_vmess",
		"https://t.me/FreeVlessVpn",
		"https://t.me/freeland8",
		"https://t.me/vmess_vless_v2rayng",
		"https://t.me/PrivateVPNs",
		"https://t.me/vmessiran",
		"https://t.me/Outline_Vpn",
		"https://t.me/vmessq",
		"https://t.me/WeePeeN",
		"https://t.me/V2rayNG3",
		"https://t.me/ShadowsocksM",
		"https://t.me/shadowsocksshop",
		"https://t.me/v2rayan",
		"https://t.me/ShadowSocks_s",
		"https://t.me/VmessProtocol",
		"https://t.me/napsternetv_config",
		"https://t.me/Easy_Free_VPN",
		"https://t.me/V2Ray_FreedomIran",
		"https://t.me/V2RAY_VMESS_free",
		"https://t.me/v2ray_for_free",
		"https://t.me/V2rayN_Free",
		"https://t.me/free4allVPN",
		"https://t.me/vpn_ocean",
		"https://t.me/configV2rayForFree",
		"https://t.me/DigiV2ray",
		"https://t.me/v2rayNG_VPN",
		"https://t.me/freev2rayssr",
		"https://t.me/v2rayn_server",
		"https://t.me/Shadowlinkserverr",
		"https://t.me/iranvpnet",
		"https://t.me/vmess_iran",
		"https://t.me/mahsaamoon1",
		"https://t.me/V2RAY_NEW",
		"https://t.me/v2RayChannel",
		"https://t.me/config_v2ray",
		"https://t.me/vpn_proxy_custom",
		"https://t.me/v2ray_custom",
		"https://t.me/VPNCUSTOMIZE",
		"https://t.me/HTTPCustomLand",
		"https://t.me/vpn_proxy_custom",
		"https://t.me/ViPVpn_v2ray",
		"https://t.me/FreeNet1500",
		"https://t.me/beta_v2ray",
		"https://t.me/vip_vpn_2022",
		"https://t.me/FOX_VPN66",
		"https://t.me/VorTexIRN",
		"https://t.me/YtTe3la",
		"https://t.me/V2RayOxygen",
		"https://t.me/Network_442",
		"https://t.me/VPN_443",
		"https://t.me/v2rayng_v",
		"https://t.me/ultrasurf_12",
		"https://t.me/frev2rayng",
		"https://t.me/frev2ray",
		"https://t.me/FreakConfig",
		"https://t.me/Awlix_ir",
		"https://t.me/v2rayngvpn",
		"https://t.me/Configforvpn01",
		"https://t.me/polproxy",
		"https://t.me/v2rayvpnchannel",
		"https://t.me/proxy_mtm",
		"https://t.me/V2Ray_FreedomIran",
		"https://t.me/v2rayfree1",
		"https://t.me/free_v2rayyy",
		"https://t.me/nx_v2ray",
		"https://t.me/nufilter",
		"https://t.me/Free_HTTPCustom",
		"https://t.me/customv2ray",
		"https://t.me/vpn_Nv1",
		"https://t.me/shopingv2ray",
		"https://t.me/v2rayng_vpnrog",
		"https://t.me/ServerNett",
		"https://t.me/MT_TEAM_IRAN",
		"https://t.me/V2ray_Team",
		"https://t.me/VpnProsecc",
		"https://t.me/ConfigsHUB",
		"https://t.me/melov2ray",
		"https://t.me/V2pedia",
		"https://t.me/FalconPolV2rayNG",
		"https://t.me/ShadowProxy66",
		"https://t.me/VPNCUSTOMIZE",
		"https://t.me/prrofile_purple",
		"https://t.me/MsV2ray",
		"https://t.me/VlessConfig",
		"https://t.me/vless_vmess",
		"https://t.me/MehradLearn",
		"https://t.me/kingofilter",
		"https://t.me/IRN_VPN",
		"https://t.me/V2raysFree",	
		"https://t.me/SvnTeam",
		"https://t.me/flyv2ray",
		"https://t.me/free1_vpn",
		"https://t.me/UnlimitedDev",
		"https://t.me/vpn_xw",
		"https://t.me/V2RayTz",
		"https://t.me/ipV2Ray",
		"https://t.me/OutlineVpnOfficial",
		"https://t.me/mehrosaboran",
		"https://t.me/mftizi",
		"https://t.me/https_config_injector",
		"https://t.me/Hope_Net",
		"https://t.me/V2rayng_Fast",
		"https://t.me/DailyV2RY",
		"https://t.me/shh_proxy",
		"https://t.me/forwardv2ray",
		"https://t.me/Lockey_vpn",
		"https://t.me/shadow_socks1",
		"https://t.me/Everyday_VPN",
		"https://t.me/V2rayCollectorDonate",
		"https://t.me/eliya_chiter0",
		"https://t.me/v2ray1_ng",
		"https://t.me/Romax_VPN",
		"https://t.me/DailyV2RY",
		"https://t.me/DeamNet_proxy",
		"https://t.me/socks5tobefree",
		"https://t.me/V_2rey",
		"https://t.me/teamkingo",
		"https://t.me/spdnet",
		"https://t.me/TunelProV2",
		"https://t.me/iran_v2ray1",
		"https://t.me/v2rayvpn2",
		"https://t.me/v2rayNgg_iran",
		"https://t.me/V2rayNG_account_free",
		"https://t.me/V2rayi_net",
		"https://t.me/vpnowl",
		"https://t.me/ghalagyann",
		"https://t.me/XpnTeam",
		"https://t.me/fhkllvjkll",
		"https://t.me/V2ray_Alpha",
		"https://t.me/IRANVPNNET",
		"https://t.me/GH_v2rayng",
		"https://t.me/VPNwedbaz",
		"https://t.me/proxyymeliii",
		"https://t.me/icv2ray",
		"https://t.me/iTV2RAY",
		"https://t.me/ArV2ray",
		"https://t.me/SEVEN_ping",
		"https://t.me/meli_proxyy",
		"https://t.me/SafeNetIR",
		"https://t.me/SafeNet_Server",
		"https://t.me/V2RAYROZ",
		"https://t.me/V2Graphy",
		"https://t.me/svn3team",
		"https://t.me/v20reyng",
		"https://t.me/s/v2rayng_org",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/flyv2ray",
		"https://t.me/s/v2ray_outlineir",
		"https://t.me/s/v2_vmess",
		"https://t.me/s/FreeVlessVpn",
		"https://t.me/s/freeland8",
		"https://t.me/s/vmess_vless_v2rayng",
		"https://t.me/s/PrivateVPNs",
		"https://t.me/s/vmessiran",
		"https://t.me/s/Outline_Vpn",
		"https://t.me/s/vmessq",
		"https://t.me/s/WeePeeN",
		"https://t.me/s/V2rayNG3",
		"https://t.me/s/ShadowsocksM",
		"https://t.me/s/shadowsocksshop",
		"https://t.me/s/v2rayan",
		"https://t.me/s/ShadowSocks_s",
		"https://t.me/s/VmessProtocol",
		"https://t.me/s/napsternetv_config",
		"https://t.me/s/Easy_Free_VPN",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/V2RAY_VMESS_free",
		"https://t.me/s/v2ray_for_free",
		"https://t.me/s/V2rayN_Free",
		"https://t.me/s/free4allVPN",
		"https://t.me/s/vpn_ocean",
		"https://t.me/s/configV2rayForFree",
		"https://t.me/s/FreeV2rays",
		"https://t.me/s/DigiV2ray",
		"https://t.me/s/v2rayNG_VPN",
		"https://t.me/s/freev2rayssr",
		"https://t.me/s/v2rayn_server",
		"https://t.me/s/Shadowlinkserverr",
		"https://t.me/s/iranvpnet",
		"https://t.me/s/vmess_iran",
		"https://t.me/s/mahsaamoon1",
		"https://t.me/s/V2RAY_NEW",
		"https://t.me/s/v2RayChannel",
		"https://t.me/s/configV2rayNG",
		"https://t.me/s/config_v2ray",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/vpnmasi",
		"https://t.me/s/v2ray_custom",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/HTTPCustomLand",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/ViPVpn_v2ray",
		"https://t.me/s/FreeNet1500",
		"https://t.me/s/v2ray_ar",
		"https://t.me/s/beta_v2ray",
		"https://t.me/s/vip_vpn_2022",
		"https://t.me/s/FOX_VPN66",
		"https://t.me/s/VorTexIRN",
		"https://t.me/s/YtTe3la",
		"https://t.me/s/V2RayOxygen",
		"https://t.me/s/Network_442",
		"https://t.me/s/VPN_443",
		"https://t.me/s/v2rayng_v",
		"https://t.me/s/ultrasurf_12",
		"https://t.me/s/iSeqaro",
		"https://t.me/s/frev2rayng",
		"https://t.me/s/frev2ray",
		"https://t.me/s/FreakConfig",
		"https://t.me/s/Awlix_ir",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/God_CONFIG",
		"https://t.me/s/Configforvpn01",
		"https://t.me/s/polproxy",
		"https://t.me/s/v2rayvpnchannel",
		"https://t.me/s/proxy_mtm",
		"https://t.me/s/vpn_ioss",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/v2rayfree1",
		"https://t.me/s/free_v2rayyy",
		"https://t.me/s/nx_v2ray",
		"https://t.me/s/nufilter",
		"https://t.me/s/Free_HTTPCustom",
		"https://t.me/s/customv2ray",
		"https://t.me/s/vpn_Nv1",
		"https://t.me/s/AliAlma_GSM",
		"https://t.me/s/reality_daily",
		"https://t.me/s/shopingv2ray",
		"https://t.me/s/v2rayng_vpnrog",
		"https://t.me/s/ServerNett",
		"https://t.me/s/MT_TEAM_IRAN",
		"https://t.me/s/V2ray_Team",
		"https://t.me/s/VpnProsecc",
		"https://t.me/s/ConfigsHUB",
		"https://t.me/s/melov2ray",
		"https://t.me/s/V2pedia",
		"https://t.me/s/FalconPolV2rayNG",
		"https://t.me/s/ShadowProxy66",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/prrofile_purple",
		"https://t.me/s/MsV2ray",
		"https://t.me/s/VlessConfig",
		"https://t.me/s/vless_vmess",
		"https://t.me/s/MehradLearn",
		"https://t.me/s/v2rayng_fa2",
		"https://t.me/s/v2rayng_org",
		"https://t.me/s/V2rayNGvpni",
		"https://t.me/s/custom_14",
		"https://t.me/s/v2rayNG_VPNN",
		"https://t.me/s/v2ray_outlineir",
		"https://t.me/s/v2_vmess",
		"https://t.me/s/FreeVlessVpn",
		"https://t.me/s/freeland8",
		"https://t.me/s/vmess_vless_v2rayng",
		"https://t.me/s/PrivateVPNs",
		"https://t.me/s/vmessiran",
		"https://t.me/s/Outline_Vpn",
		"https://t.me/s/vmessq",
		"https://t.me/s/WeePeeN",
		"https://t.me/s/V2rayNG3",
		"https://t.me/s/ShadowsocksM",
		"https://t.me/s/shadowsocksshop",
		"https://t.me/s/v2rayan",
		"https://t.me/s/ShadowSocks_s",
		"https://t.me/s/VmessProtocol",
		"https://t.me/s/napsternetv_config",
		"https://t.me/s/Easy_Free_VPN",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/V2RAY_VMESS_free",
		"https://t.me/s/v2ray_for_free",
		"https://t.me/s/V2rayN_Free",
		"https://t.me/s/free4allVPN",
		"https://t.me/s/vpn_ocean",
		"https://t.me/s/configV2rayForFree",
		"https://t.me/s/FreeV2rays",
		"https://t.me/s/DigiV2ray",
		"https://t.me/s/v2rayNG_VPN",
		"https://t.me/s/freev2rayssr",
		"https://t.me/s/v2rayn_server",
		"https://t.me/s/Shadowlinkserverr",
		"https://t.me/s/iranvpnet",
		"https://t.me/s/vmess_iran",
		"https://t.me/s/mahsaamoon1",
		"https://t.me/s/V2RAY_NEW",
		"https://t.me/s/v2RayChannel",
		"https://t.me/s/configV2rayNG",
		"https://t.me/s/config_v2ray",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/vpnmasi",
		"https://t.me/s/v2ray_custom",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/HTTPCustomLand",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/ViPVpn_v2ray",
		"https://t.me/s/FreeNet1500",
		"https://t.me/s/v2ray_ar",
		"https://t.me/s/beta_v2ray",
		"https://t.me/s/vip_vpn_2022",
		"https://t.me/s/FOX_VPN66",
		"https://t.me/s/VorTexIRN",
		"https://t.me/s/YtTe3la",
		"https://t.me/s/V2RayOxygen",
		"https://t.me/s/Network_442",
		"https://t.me/s/VPN_443",
		"https://t.me/s/v2rayng_v",
		"https://t.me/s/ultrasurf_12",
		"https://t.me/s/iSeqaro",
		"https://t.me/s/frev2rayng",
		"https://t.me/s/frev2ray",
		"https://t.me/s/FreakConfig",
		"https://t.me/s/Awlix_ir",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/God_CONFIG",
		"https://t.me/s/Configforvpn01",
		"https://t.me/s/polproxy",
		"https://t.me/s/v2rayvpnchannel",
		"https://t.me/s/proxy_mtm",
		"https://t.me/s/vpn_ioss",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/v2rayfree1",
		"https://t.me/s/free_v2rayyy",
		"https://t.me/s/nx_v2ray",
		"https://t.me/s/nufilter",
		"https://t.me/s/Free_HTTPCustom",
		"https://t.me/s/customv2ray",
		"https://t.me/s/PewezaVPN",
		"https://t.me/s/v2rayng_org",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/flyv2ray",
		"https://t.me/s/v2ray_outlineir",
		"https://t.me/s/v2_vmess",
		"https://t.me/s/FreeVlessVpn",
		"https://t.me/s/freeland8",
		"https://t.me/s/vmess_vless_v2rayng",
		"https://t.me/s/PrivateVPNs",
		"https://t.me/s/vmessiran",
		"https://t.me/s/Outline_Vpn",
		"https://t.me/s/vmessq",
		"https://t.me/s/WeePeeN",
		"https://t.me/s/V2rayNG3",
		"https://t.me/s/ShadowsocksM",
		"https://t.me/s/shadowsocksshop",
		"https://t.me/s/v2rayan",
		"https://t.me/s/ShadowSocks_s",
		"https://t.me/s/VmessProtocol",
		"https://t.me/s/napsternetv_config",
		"https://t.me/s/Easy_Free_VPN",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/V2RAY_VMESS_free",
		"https://t.me/s/v2ray_for_free",
		"https://t.me/s/V2rayN_Free",
		"https://t.me/s/free4allVPN",
		"https://t.me/s/vpn_ocean",
		"https://t.me/s/configV2rayForFree",
		"https://t.me/s/FreeV2rays",
		"https://t.me/s/DigiV2ray",
		"https://t.me/s/v2rayNG_VPN",
		"https://t.me/s/freev2rayssr",
		"https://t.me/s/v2rayn_server",
		"https://t.me/s/Shadowlinkserverr",
		"https://t.me/s/iranvpnet",
		"https://t.me/s/vmess_iran",
		"https://t.me/s/mahsaamoon1",
		"https://t.me/s/V2RAY_NEW",
		"https://t.me/s/v2RayChannel",
		"https://t.me/s/configV2rayNG",
		"https://t.me/s/config_v2ray",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/vpnmasi",
		"https://t.me/s/v2ray_custom",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/HTTPCustomLand",
		"https://t.me/s/vpn_proxy_custom",
		"https://t.me/s/ViPVpn_v2ray",
		"https://t.me/s/FreeNet1500",
		"https://t.me/s/v2ray_ar",
		"https://t.me/s/beta_v2ray",
		"https://t.me/s/vip_vpn_2022",
		"https://t.me/s/FOX_VPN66",
		"https://t.me/s/VorTexIRN",
		"https://t.me/s/YtTe3la",
		"https://t.me/s/V2RayOxygen",
		"https://t.me/s/Network_442",
		"https://t.me/s/VPN_443",
		"https://t.me/s/v2rayng_v",
		"https://t.me/s/ultrasurf_12",
		"https://t.me/s/iSeqaro",
		"https://t.me/s/frev2rayng",
		"https://t.me/s/frev2ray",
		"https://t.me/s/FreakConfig",
		"https://t.me/s/Awlix_ir",
		"https://t.me/s/v2rayngvpn",
		"https://t.me/s/God_CONFIG",
		"https://t.me/s/Configforvpn01",
		"https://t.me/s/polproxy",
		"https://t.me/s/v2rayvpnchannel",
		"https://t.me/s/proxy_mtm",
		"https://t.me/s/vpn_ioss",
		"https://t.me/s/V2Ray_FreedomIran",
		"https://t.me/s/v2rayfree1",
		"https://t.me/s/free_v2rayyy",
		"https://t.me/s/nx_v2ray",
		"https://t.me/s/nufilter",
		"https://t.me/s/Free_HTTPCustom",
		"https://t.me/s/customv2ray",
		"https://t.me/s/vpn_Nv1",
		"https://t.me/s/AliAlma_GSM",
		"https://t.me/s/reality_daily",
		"https://t.me/s/shopingv2ray",
		"https://t.me/s/v2rayng_vpnrog",
		"https://t.me/s/ServerNett",
		"https://t.me/s/MT_TEAM_IRAN",
		"https://t.me/s/V2ray_Team",
		"https://t.me/s/VpnProsecc",
		"https://t.me/s/ConfigsHUB",
		"https://t.me/s/melov2ray",
		"https://t.me/s/V2pedia",
		"https://t.me/s/FalconPolV2rayNG",
		"https://t.me/s/ShadowProxy66",
		"https://t.me/s/VPNCUSTOMIZE",
		"https://t.me/s/prrofile_purple",
		"https://t.me/s/MsV2ray",
		"https://t.me/s/VlessConfig",
		"https://t.me/s/vless_vmess",
		"https://t.me/s/MehradLearn",
		"https://t.me/s/kingofilter",
		"https://t.me/s/IRN_VPN",
		"https://t.me/s/V2raysFree",
		"https://t.me/s/SvnTeam",
		"https://t.me/s/flyv2ray",
		"https://t.me/s/free1_vpn",
		"https://t.me/s/UnlimitedDev",
		"https://t.me/s/vpn_xw",
		"https://t.me/s/V2RayTz",
		"https://t.me/s/ipV2Ray",
		"https://t.me/s/OutlineVpnOfficial",
		"https://t.me/s/mehrosaboran",
		"https://t.me/s/mftizi",
		"https://t.me/s/https_config_injector",
		"https://t.me/s/Hope_Net",
		"https://t.me/s/V2rayng_Fast",
		"https://t.me/s/DailyV2RY",
		"https://t.me/s/shh_proxy",
		"https://t.me/s/forwardv2ray",
		"https://t.me/s/Lockey_vpn",
		

	}

	configs := map[string]string{
		"":     "",
		"vmess":  "",
		"trojan": "",
		"vless":  "",
		"mixed":  "",
	}

	myregex := map[string]string{
		"":     `(.{3})ss:\/\/`,
		"vmess":  `vmess:\/\/`,
		"trojan": `trojan:\/\/`,
		"vless":  `vless:\/\/`,
	}






	
	//protocol := ""
	for i := 0; i < len(channels); i++ {
		all_messages := false
		if strings.Contains(channels[i], "{all_messages}") {
			all_messages = true
			channels[i] = strings.Split(channels[i], "{all_messages}")[0]
		}

		req, err := http.NewRequest("GET", channels[i], nil)
		if err != nil {
			log.Fatalf("Error When requesting to: %s Error : %s", channels[i], err)
		}

		resp, err1 := client.Do(req)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		messages := doc.Find(".tgme_widget_message_wrap").Length()
		link, exist := doc.Find(".tgme_widget_message_wrap .js-widget_message").Last().Attr("data-post")
		if messages < 20 && exist {
			number := strings.Split(link, "/")[1]
			fmt.Println(number)

			doc = GetMessages(20, doc, number, channels[i])
		}

		if all_messages {
			fmt.Println(doc.Find(".js-widget_message_wrap").Length())
			doc.Find(".tgme_widget_message_text").Each(func(j int, s *goquery.Selection) {
				// For each item found, get the band and title
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for _, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							return "\n" + match
						})
					}
					for proto, _ := range configs {
						if strings.Contains(lines[a], proto) {
							configs["mixed"] += "\n" + lines[a] + "\n"
						}
					}
				}

			})
		} else {
			doc.Find("code,pre").Each(func(j int, s *goquery.Selection) {
				// For each item found, get the band and title
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for proto_regex, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							if proto_regex == "" {
								if match[:3] == "vme" {
									return "\n" + match
								} else if match[:3] == "vle" {
									return "\n" + match
								} else {
									return "\n" + match
								}
							} else {
								return "\n" + match
							}
						})

						if len(strings.Split(lines[a], "#")) > 1 {
							myconfigs := strings.Split(lines[a], "#")
							for i := 0; i < len(myconfigs); i++ {
								if myconfigs[i] != "" {
									re := regexp.MustCompile(regex_value)
								myconfigs[i] = strings.ReplaceAll(myconfigs[i], "%40IRAn_V2Ray1" , "اتصال")
								
										

									
									match := re.FindStringSubmatch(myconfigs[i])
									if len(match) >= 1 {
										if proto_regex == "" {
											if match[1][:3] == "vme" {
												configs[""] += "\n" + myconfigs[i] + "\n"
											} else if match[1][:3] == "vle" {
												configs[""] += "\n" + myconfigs[i] + "\n"
											} else {
												configs[""] += "\n" + myconfigs[i][3:] + "\n"
											}
										} else {
											configs[proto_regex] += "\n" + myconfigs[i] + "\n"
										}
									}

								}

							}
						}
					}
				}
			})
		}

	}


	for proto, configcontent := range configs {
		// 		reverse mode :
		// 		lines := strings.Split(configcontent, "\n")
		// 		reversed := reverse(lines)
		// 		WriteToFile(strings.Join(reversed, "\n"), proto+"_iran.txt")
		// 		simple mode :
		WriteToFile(RemoveDuplicate(configcontent), proto+"update")
	}

}




func WriteToFile(fileContent string, filePath string) {

	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		// If the file exists, clear its content
		err = ioutil.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Println("Error clearing file:", err)
			return
		}
	} else if os.IsNotExist(err) {
		// If the file does not exist, create it
		_, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	} else {
		// If there was some other error, print it and return
		fmt.Println("Error checking file:", err)
		return
	}

	// Write the new content to the file
	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func load_more(link string) *goquery.Document {
	req, _ := http.NewRequest("GET", link, nil)
	fmt.Println(link)
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return doc
}






func GetMessages(length int, doc *goquery.Document, number string, channel string) *goquery.Document {
	x := load_more(channel + "?before=" + number)

	html2, _ := x.Html()
	reader2 := strings.NewReader(html2)
	doc2, _ := goquery.NewDocumentFromReader(reader2)

	// _, exist := doc.Find(".js-messages_more").Attr("href")
	doc.Find("body").AppendSelection(doc2.Find("body").Children())

	newDoc := goquery.NewDocumentFromNode(doc.Selection.Nodes[0])
	// fmt.Println(newDoc.Find(".js-messages_more").Attr("href"))
	messages := newDoc.Find(".js-widget_message_wrap").Length()

	fmt.Println(messages)
	if messages > length {
		return newDoc
	} else {
		num, _ := strconv.Atoi(number)
		n := num - 21
		if n > 0 {
			ns := strconv.Itoa(n)
			GetMessages(length, newDoc, ns, channel)
		} else {
			return newDoc
		}
	}

	return newDoc
}

func reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}

	return lines
}

func RemoveDuplicate(config string) string {
	lines := strings.Split(config, "\n")

	// Use a map to keep track of unique lines
	uniqueLines := make(map[string]bool)

	// Loop over lines and add unique lines to map
	for _, line := range lines {
		if len(line) > 0 {
			uniqueLines[line] = true
		}
	}

	// Join unique lines into a string
	uniqueString := strings.Join(getKeys(uniqueLines), "\n")

	return uniqueString
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
