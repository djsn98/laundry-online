import React from 'react'
import { Dimensions, Image, ImageBackground, ScrollView, StyleSheet, Text, View } from 'react-native'
import { ImageHeader, Logo } from '../../assets'
import { ButtonIcon, Saldo, PesananAktif } from '../../components'
import { WARNA_ABU_ABU } from '../../utils/constant'

const Home = () => {
    return (
        <View style={styles.page}>
            <ScrollView showsVerticalScrollIndicator={false}>
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
                        <ButtonIcon title="Kiloan" type="layanan" />
                        <ButtonIcon title="Satuan" type="layanan" />
                        <ButtonIcon title="VIP" type="layanan" />
                        <ButtonIcon title="Karpet" type="layanan" />
                        <ButtonIcon title="Setrika Saja" type="layanan" />
                        <ButtonIcon title="Express" type="layanan" />
                    </View>
                </View>
                <View style={styles.pesananAktif}>
                    <Text style={styles.label}>Pesanan Aktif</Text>
                    <PesananAktif orderId="0001212" status="Sudah Selesai"/>
                    <PesananAktif orderId="0001213" status="Masih Dicuci"/>
                    <PesananAktif orderId="0001214" status="Sudah Selesai"/>
                    <PesananAktif orderId="0001215" status="Sudah Selesai"/>
                </View>
            </ScrollView>
        </View>
    )
}

export default Home

const windowWidth = Dimensions.get('window').width
const windowHeight = Dimensions.get('window').height

const styles = StyleSheet.create({
    page: {
        flex: 1,
        backgroundColor: 'white'
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
        marginTop: 10,
        flexWrap: 'wrap',
    },
    pesananAktif: {
        paddingVertical: 10,
        paddingHorizontal: 30,
        backgroundColor: WARNA_ABU_ABU,
        flex: 1,
        borderTopRightRadius: 20,
        borderTopLeftRadius: 20,
    }
})
