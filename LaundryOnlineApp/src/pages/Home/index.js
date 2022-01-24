import React from 'react'
import { Dimensions, Image, ImageBackground, StyleSheet, Text, View } from 'react-native'
import { ImageHeader, Logo } from '../../assets'
import { ButtonIcon, Saldo } from '../../components'

const Home = () => {
    return (
        <View style={styles.page}>
            <ImageBackground source={ImageHeader} style={styles.header}>
                <Image source={Logo} style={styles.logo} />
                <View style={styles.hello}>
                    <Text style={styles.selamat}>Selamat Datang,</Text>
                    <Text style={styles.username}>Dennis Jason</Text>
                </View>
            </ImageBackground>
            <Saldo />
            <View style={styles.layanan}>
                <Text style={styles.label}>Layanan Kami</Text>
                <View style={styles.iconLayanan}>
                    <ButtonIcon title="Kiloan" type="layanan"/>
                    <ButtonIcon title="Satuan" type="layanan"/>
                    <ButtonIcon title="VIP" type="layanan"/>
                    <ButtonIcon title="Karpet" type="layanan"/>
                    <ButtonIcon title="Setrika Saja" type="layanan"/>
                    <ButtonIcon title="Express" type="layanan"/>
                </View>
            </View>
        </View>
    )
}

export default Home

const windowWidth = Dimensions.get('window').width
const windowHeight = Dimensions.get('window').height

const styles = StyleSheet.create({
    page: {
        flex: 1
    },
    header: {
        width: windowWidth,
        height: windowHeight * 0.3,
        paddingHorizontal: 30,
        paddingTop: 10
    },
    logo: {
        width: windowWidth * 0.25,
        height: windowHeight * 0.06
    },
    hello: {
        marginTop: windowHeight * 0.030
    },
    selamat: {
        fontSize: 22,
        fontFamily: 'TitilliumWeb-Regular'
    },
    username: {
        fontSize: 20,
        fontFamily: 'TitilliumWeb-Bold'
    },
    layanan: {
        paddingLeft: 30,
        paddingTop: 15,
    },
    label: {
        fontSize: 18,
        fontFamily: 'TitilliumWeb-Bold',
    },
    iconLayanan: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginTop: 12,
        flexWrap: 'wrap',
    }
})
